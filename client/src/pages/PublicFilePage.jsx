import { useEffect, useState } from 'react';
import { filesApi } from '../api/files';

export default function PublicFilePage() {
  const fileId = window.location.pathname.replace(/^\/f\//, '').split('/')[0];
  const [file, setFile] = useState(null);
  const [downloadUrl, setDownloadUrl] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!fileId) {
      setError('Invalid file link');
      setLoading(false);
      return;
    }

    let cancelled = false;

    async function load() {
      try {
        const [detailsRes, downloadRes] = await Promise.all([
          filesApi.getDetails(fileId),
          filesApi.getDownloadUrl(fileId),
        ]);
        if (cancelled) return;
        setFile(detailsRes);
        const url = downloadRes?.download_url ?? downloadRes?.url;
        if (url) setDownloadUrl(url);
      } catch (err) {
        if (!cancelled) {
          setError(err.message || 'File not found or not public');
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    }

    load();
    return () => { cancelled = true; };
  }, [fileId]);

  if (loading) {
    return (
      <div className="min-h-screen bg-gv-dark flex items-center justify-center text-gray-400">
        <div className="animate-pulse">Loading…</div>
      </div>
    );
  }

  if (error || !file) {
    return (
      <div className="min-h-screen bg-gv-dark flex items-center justify-center text-red-400">
        <div className="text-center">
          <p className="text-lg font-medium">{error || 'File not found'}</p>
          <a href="/" className="mt-4 inline-block text-blue-400 hover:underline">Go back</a>
        </div>
      </div>
    );
  }

  const name = file.name || 'File';
  const mime = (file.mime_type || '').toLowerCase();
  const isImage = mime.startsWith('image/');
  const isPdf = mime === 'application/pdf';

  return (
    <div className="min-h-screen bg-gv-dark text-gray-200 p-6">
      <div className="max-w-4xl mx-auto">
        <div className="mb-4 flex items-center justify-between gap-4 flex-wrap">
          <h1 className="text-xl font-semibold text-white truncate" title={name}>{name}</h1>
          {downloadUrl && (
            <a
              href={downloadUrl}
              download={name}
              target="_blank"
              rel="noopener noreferrer"
              className="shrink-0 px-4 py-2 rounded-lg bg-blue-600 hover:bg-blue-500 text-white text-sm font-medium transition-colors"
            >
              Download
            </a>
          )}
        </div>

        <div className="bg-[#161b22] border border-[#30363d] rounded-xl overflow-hidden">
          {isImage && downloadUrl && (
            <div className="p-4 flex justify-center bg-gv-dark">
              <img src={downloadUrl} alt={name} className="max-w-full max-h-[80vh] object-contain" />
            </div>
          )}
          {isPdf && downloadUrl && (
            <div className="w-full h-[80vh] flex justify-center bg-gv-dark">
              <iframe
                src={downloadUrl}
                title={name}
                className="w-full flex-1 border-0 rounded"
              />
            </div>
          )}
          {!isImage && !isPdf && (
            <div className="p-8 text-center text-gray-400">
              <p className="mb-2">{file.mime_type || 'File'}</p>
              {downloadUrl ? (
                <a
                  href={downloadUrl}
                  download={name}
                  className="text-blue-400 hover:underline"
                >
                  Download file
                </a>
              ) : (
                <p>Download not available</p>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

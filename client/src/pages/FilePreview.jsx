import { useEffect, useState } from 'react';
import { filesApi } from '../api/files';

export default function FilePreview() {
  const fileId = window.location.pathname.replace(/^\/f\//, '').split('/')[0];
  const [file, setFile] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);
  const [downloading, setDownloading] = useState(false);

  useEffect(() => {
    if (!fileId) {
      setError('Invalid file link');
      setLoading(false);
      return;
    }

    let cancelled = false;

    async function load() {
      try {
        const detailsRes = await filesApi.getDetails(fileId);
        if (cancelled) return;
        setFile(detailsRes);
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

  const handleDownload = async () => {
    if (downloading) return;
    setDownloading(true);
    try {
      const downloadRes = await filesApi.getDownloadUrl(fileId);
      const url = downloadRes?.download_url ?? downloadRes?.url;
      if (url) {
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', file?.name || 'file');
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      } else {
        alert('Download URL not available');
      }
    } catch (err) {
      alert('Download failed: ' + err.message);
    } finally {
      setDownloading(false);
    }
  };

  const formatSize = (bytes) => {
    if (!bytes) return 'Unknown size';
    if (bytes < 1024) return bytes + " B";
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(2) + " KB";
    if (bytes < 1024 * 1024 * 1024) return (bytes / (1024 * 1024)).toFixed(2) + " MB";
    return (bytes / (1024 * 1024 * 1024)).toFixed(2) + " GB";
  };

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
  const size = formatSize(file.size_bytes);

  return (
    <div className="min-h-screen bg-gv-dark text-gray-200 p-6">
      <div className="max-w-4xl mx-auto">
        <div className="mb-6">
          <h1 className="text-2xl font-semibold text-white mb-2">{name}</h1>
          <p className="text-sm text-gray-400">Size: {size}</p>
        </div>

        <div className="bg-[#161b22] border border-[#30363d] rounded-xl p-6">
          <div className="text-center">
            <button
              onClick={handleDownload}
              disabled={downloading}
              className="inline-block px-6 py-3 rounded-lg bg-blue-600 hover:bg-blue-500 disabled:opacity-50 disabled:cursor-not-allowed text-white text-sm font-medium transition-colors"
            >
              {downloading ? 'Preparing download...' : 'Download'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

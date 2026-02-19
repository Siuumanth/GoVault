import { useEffect, useState, useCallback } from 'react';
import Sidebar from '../components/layout/Sidebar';
import Navbar from '../components/layout/Navbar';
import FileCard from '../components/FileCard';
import ShareModal from '../components/ShareModal';
import UploadView from '../components/upload/UploadView'; // New View
import { filesApi } from '../api/files';
import { useUpload } from '../hooks/useUpload';
import { useAuth } from '../context/AuthContext';

export default function Dashboard() {
  const [activeTab, setActiveTab] = useState('owned');
  const [files, setFiles] = useState([]);
  const [selectedFileForShare, setSelectedFileForShare] = useState(null);
  
  const { uploadFile, progress, isUploading, logs } = useUpload();
  const { logout } = useAuth();

  const loadData = useCallback(async () => {
    if (activeTab === 'upload') return; // Don't fetch files if on upload page
    try {
      let data;
      if (activeTab === 'owned') data = await filesApi.getOwned();
      else if (activeTab === 'shared') data = await filesApi.getShared();
      else if (activeTab === 'shortcuts') data = await filesApi.getShortcuts();
      setFiles(data?.files || []);
    } catch (err) {
      setFiles([]);
    }
  }, [activeTab]);

  useEffect(() => { loadData(); }, [loadData]);

  const handleUploadProcess = async (file) => {
    const success = await uploadFile(file);
    if (success) {
      // Stay on upload page to show "Upload Done" message
      // User can manually switch back to 'owned' to see the file
    }
  };

  const handleDownload = async (fileId, fileName) => {
    const data = await filesApi.getDownloadUrl(fileId);
    if (data?.download_url) {
      const link = document.createElement('a');
      link.href = data.download_url;
      link.setAttribute('download', fileName);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    }
  };

  return (
    <div className="flex h-screen bg-[#0d1117] text-gray-200 overflow-hidden">
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="flex-1 flex flex-col min-w-0">
        <Navbar />

        {/* Sub-Header */}
        <header className="h-14 border-b border-[#30363d] bg-[#0d1117] flex items-center justify-between px-8 shrink-0">
          <h2 className="text-white font-semibold text-lg capitalize">{activeTab}</h2>
          {/* Quick upload button removed from header since we have a dedicated page now */}
        </header>

        <section className="flex-1 overflow-y-auto p-8">
          
          {/* Conditional Rendering based on Tab */}
          {activeTab === 'upload' ? (
            <UploadView 
              onUpload={handleUploadProcess} 
              progress={progress} 
              isUploading={isUploading} 
              logs={logs} 
            />
          ) : (
            <>
              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6">
                {files.map(file => (
                  <FileCard 
                    key={file.file_id} file={file} 
                    onDownload={handleDownload}
                    onDelete={() => filesApi.delete(file.file_id).then(loadData)}
                    onRename={() => {/* logic */}}
                    onShare={() => setSelectedFileForShare(file)}
                  />
                ))}
              </div>
              {files.length === 0 && (
                <div className="flex flex-col items-center justify-center py-40 opacity-20 text-6xl">📂</div>
              )}
            </>
          )}
        </section>
      </div>

      {selectedFileForShare && (
        <ShareModal file={selectedFileForShare} onClose={() => setSelectedFileForShare(null)} />
      )}
    </div>
  );
}
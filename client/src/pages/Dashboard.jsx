import { useEffect, useState, useCallback } from 'react';
import Sidebar from '../components/layout/Sidebar';
import Navbar from '../components/layout/Navbar';
import FileCard from '../components/FileCard';
import ShareModal from '../components/ShareModal';
import RenameModal from '../components/RenameModal';
import UploadView from '../components/upload/UploadView';
import { filesApi } from '../api/files';
import { useUpload } from '../hooks/useUpload';

export default function Dashboard() {
  // --- State ---
  const [activeTab, setActiveTab] = useState('owned');
  const [files, setFiles] = useState([]);
  const [shortcutIds, setShortcutIds] = useState(new Set());
  const [publicFileIds, setPublicFileIds] = useState(new Set());
  
  // Modal State
  const [selectedFile, setSelectedFile] = useState(null);
  const [modalType, setModalType] = useState(null); // 'share' or 'rename'

  // Hooks
  const { uploadFile, progress, isUploading, logs } = useUpload();

  // --- Data Loading ---
  const loadData = useCallback(async () => {
    if (activeTab === 'upload') return; 

    try {
      let data;
      if (activeTab === 'owned') data = await filesApi.getOwned();
      else if (activeTab === 'shared') data = await filesApi.getShared();
      else if (activeTab === 'shortcuts') data = await filesApi.getShortcuts();
      
      let fileList = data?.files || data || [];
      
      // Sort: Newest First
      fileList.sort((a, b) => {
        const dateA = new Date(a.created_at || 0);
        const dateB = new Date(b.created_at || 0);
        return dateB - dateA;
      });
      
      setFiles(fileList);

      // Handle Shortcut Status
      const shortcutsData = await filesApi.getShortcuts();
      const ids = (shortcutsData?.files || []).map(f => f.file_id);
      setShortcutIds(new Set(ids));

      // Handle Public Status
      const publicIds = new Set();
      fileList.forEach(f => { if (f.is_public) publicIds.add(f.file_id); });
      setPublicFileIds(publicIds);

    } catch (err) {
      console.error("Fetch failed:", err);
      setFiles([]);
    }
  }, [activeTab]);

  useEffect(() => { loadData(); }, [loadData]);

  // --- Handlers ---
  const handleDownload = async (fileId, fileName) => {
    try {
      const data = await filesApi.getDownloadUrl(fileId);
      if (data?.download_url) {
        const link = document.createElement('a');
        link.href = data.download_url;
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      }
    } catch (err) { alert("Download failed"); }
  };

  const handleTogglePublic = async (file) => {
    try {
      const isPublic = publicFileIds.has(file.file_id);
      if (isPublic) {
        await filesApi.removePublic(file.file_id);
        setPublicFileIds(prev => {
          const next = new Set(prev);
          next.delete(file.file_id);
          return next;
        });
        alert('Public access removed successfully');
      } else {
        const result = await filesApi.makePublic(file.file_id);
        setPublicFileIds(prev => new Set(prev).add(file.file_id));
        if (result?.public_url) {
          alert(`File is now public!\nPublic URL: ${result.public_url}`);
        } else {
          alert('File is now public!');
        }
      }
      loadData();
    } catch (err) {
      alert("Public access operation failed: " + err.message);
    }
  };

  const openModal = (file, type) => {
    setSelectedFile(file);
    setModalType(type);
  };

  const closeModal = () => {
    setSelectedFile(null);
    setModalType(null);
  };

  // Group files by date
  const groupFilesByDate = (fileList) => {
    const groups = {};
    fileList.forEach(file => {
      const date = new Date(file.created_at || file.updated_at || 0);
      const dateKey = date.toDateString(); // Use this for grouping
      if (!groups[dateKey]) {
        groups[dateKey] = [];
      }
      groups[dateKey].push(file);
    });
    return groups;
  };

  // Format date as "20 February 2026"
  const formatDateHeader = (dateString) => {
    const date = new Date(dateString);
    const day = date.getDate();
    const month = date.toLocaleDateString('en-US', { month: 'long' });
    const year = date.getFullYear();
    return `${day} ${month} ${year}`;
  };

  return (
    <div className="flex h-screen bg-gv-dark text-gray-200 overflow-hidden">
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="flex-1 flex flex-col min-w-0">
        <Navbar />

        <header className="h-14 border-b border-[#30363d] bg-[#0d1117] flex items-center justify-between px-8 shrink-0">
          <h2 className="text-white font-semibold text-lg capitalize">{activeTab}</h2>
        </header>

        <section className="flex-1 overflow-y-auto p-8 custom-scrollbar">
          {activeTab === 'upload' ? (
            <UploadView 
              onUpload={uploadFile} 
              progress={progress} 
              isUploading={isUploading} 
              logs={logs} 
            />
          ) : (
            <>
              {/* Background Upload Progress */}
              {isUploading && (
                <div className="mb-8 p-4 bg-blue-600/5 border border-blue-500/20 rounded-xl">
                  <div className="flex justify-between text-xs mb-2">
                    <span className="text-blue-400 font-bold uppercase tracking-widest">Uploading...</span>
                    <span className="text-blue-400 font-mono">{progress}%</span>
                  </div>
                  <div className="w-full bg-[#161b22] h-1 rounded-full overflow-hidden">
                    <div className="bg-blue-500 h-full transition-all" style={{ width: `${progress}%` }} />
                  </div>
                </div>
              )}

              {/* Files grouped by date */}
              {(() => {
                const grouped = groupFilesByDate(files);
                const sortedDates = Object.keys(grouped).sort((a, b) => {
                  return new Date(b) - new Date(a); // Recent to old
                });

                return sortedDates.map(dateKey => (
                  <div key={dateKey} className="mb-8">
                    <h3 className="text-sm font-semibold text-gray-400 mb-4 uppercase tracking-wider">
                      {formatDateHeader(dateKey)}
                    </h3>
                    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-5">
                      {grouped[dateKey].map(file => (
                        <FileCard 
                          key={file.file_id} 
                          file={file} 
                          isShortcut={shortcutIds.has(file.file_id)}
                          isPublic={publicFileIds.has(file.file_id)}
                          onDownload={handleDownload}
                          onDelete={(id) => filesApi.delete(id).then(loadData)}
                          onShare={(f) => openModal(f, 'share')}
                          onRename={(f) => openModal(f, 'rename')}
                          onShortcut={(id) => filesApi.addShortcut(id).then(loadData)}
                          onTogglePublic={handleTogglePublic}
                        />
                      ))}
                    </div>
                  </div>
                ));
              })()}

              {files.length === 0 && (
                <div className="flex flex-col items-center justify-center py-40 opacity-40">
                  <div className="text-7xl mb-6">📁</div>
                  <p className="text-xl font-medium">No files found</p>
                </div>
              )}
            </>
          )}
        </section>
      </div>

      {/* Modals */}
      {selectedFile && modalType === 'share' && (
        <ShareModal file={selectedFile} onClose={closeModal} />
      )}
      
      {selectedFile && modalType === 'rename' && (
        <RenameModal 
          file={selectedFile} 
          onClose={closeModal}
          onSuccess={loadData}
        />
      )}
    </div>
  );
}
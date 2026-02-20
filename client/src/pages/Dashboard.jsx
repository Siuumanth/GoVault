import { useEffect, useState, useCallback } from 'react';
import Sidebar from '../components/layout/Sidebar';
import Navbar from '../components/layout/Navbar';
import FileCard from '../components/FileCard';
import ShareModal from '../components/ShareModal';
import RenameModal from '../components/RenameModal';
import UploadView from '../components/upload/UploadView';
import { filesApi } from '../api/files';
import { useUpload } from '../hooks/useUpload';
import { useAuth } from '../context/AuthContext';

export default function Dashboard() {
  // --- 1. State ---
  const [activeTab, setActiveTab] = useState('owned');
  const [files, setFiles] = useState([]);
  const [shortcutIds, setShortcutIds] = useState(new Set());
  const [publicFileIds, setPublicFileIds] = useState(new Set());
  
  // Modal States
  const [selectedFile, setSelectedFile] = useState(null);
  const [modalType, setModalType] = useState(null); // 'share' or 'rename'

  // --- 2. Hooks ---
  const { uploadFile, progress, isUploading, logs } = useUpload();
  const { logout } = useAuth();

  // --- 3. Data Loading ---
  const loadData = useCallback(async () => {
    if (activeTab === 'upload') return; 

    try {
      let data;
      if (activeTab === 'owned') data = await filesApi.getOwned();
      else if (activeTab === 'shared') data = await filesApi.getShared();
      else if (activeTab === 'shortcuts') data = await filesApi.getShortcuts();
      
      let fileList = data?.files || [];
      
      // Sort: Newest First
      fileList.sort((a, b) => {
        const dateA = new Date(a.created_at || 0);
        const dateB = new Date(b.created_at || 0);
        return dateB - dateA;
      });
      
      setFiles(fileList);

      // Handle Shortcut Status
      if (activeTab === 'shortcuts') {
        setShortcutIds(new Set(fileList.map(f => f.file_id)));
      } else {
        const shortcutsData = await filesApi.getShortcuts();
        const ids = (shortcutsData?.files || []).map(f => f.file_id);
        setShortcutIds(new Set(ids));
      }

      // Handle Public Status
      const publicIds = new Set();
      fileList.forEach(f => { if (f.is_public) publicIds.add(f.file_id); });
      setPublicFileIds(publicIds);

    } catch (err) {
      console.error("Failed to load dashboard data:", err);
      setFiles([]);
    }
  }, [activeTab]);

  useEffect(() => { loadData(); }, [loadData]);

  // --- 4. Handlers ---
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

  const handleShortcut = async (fileId) => {
    try {
      shortcutIds.has(fileId) 
        ? await filesApi.removeShortcut(fileId) 
        : await filesApi.addShortcut(fileId);
      loadData();
    } catch (err) { alert("Shortcut failed"); }
  };

  const handleTogglePublic = async (file) => {
    try {
      publicFileIds.has(file.file_id) 
        ? await filesApi.removePublic(file.file_id) 
        : await filesApi.makePublic(file.file_id);
      loadData();
    } catch (err) { alert("Visibility change failed"); }
  };

  const closeModal = () => {
    setSelectedFile(null);
    setModalType(null);
  };

  return (
    <div className="flex h-screen bg-[#0d1117] text-gray-200 overflow-hidden">
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="flex-1 flex flex-col min-w-0">
        <Navbar />

<<<<<<< HEAD
        {/* Sub-Header */}
        <header className="h-14 border-b border-[#30363d] bg-gv-dark flex items-center justify-between px-8 shrink-0">
=======
        <header className="h-14 border-b border-[#30363d] bg-[#0d1117] flex items-center justify-between px-8 shrink-0">
>>>>>>> client
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
              {/* Progress bar shown even on Vault page if an upload is active */}
              {isUploading && (
                <div className="mb-8 p-4 bg-blue-600/5 border border-blue-500/20 rounded-xl">
                  <div className="flex justify-between text-xs mb-2">
                    <span className="text-blue-400 font-bold uppercase tracking-widest">Uploading in Background...</span>
                    <span className="text-blue-400 font-mono">{progress}%</span>
                  </div>
                  <div className="w-full bg-[#161b22] h-1 rounded-full overflow-hidden">
                    <div className="bg-blue-500 h-full transition-all" style={{ width: `${progress}%` }} />
                  </div>
                </div>
              )}

              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-5">
                {files.map(file => (
                  <FileCard 
                    key={file.file_id} 
                    file={file} 
                    isShortcut={shortcutIds.has(file.file_id)}
                    isPublic={publicFileIds.has(file.file_id)}
                    onDownload={handleDownload}
                    onDelete={(id) => filesApi.delete(id).then(loadData)}
                    onShare={(f) => { setSelectedFile(f); setModalType('share'); }}
                    onRename={(f) => { setSelectedFile(f); setModalType('rename'); }}
                    onShortcut={handleShortcut}
                    onTogglePublic={handleTogglePublic}
                  />
                ))}
              </div>

              {files.length === 0 && (
                <div className="flex flex-col items-center justify-center py-40 opacity-40">
                  <div className="text-7xl mb-6">📁</div>
                  <p className="text-xl font-medium">No files here yet</p>
                </div>
              )}
            </>
          )}
        </section>
      </div>
<<<<<<< HEAD
=======

      {/* --- Unified Modals --- */}
      {selectedFile && modalType === 'share' && (
        <ShareModal file={selectedFile} onClose={closeModal} />
      )}
      
      {selectedFile && modalType === 'rename' && (
        <RenameModal 
          file={selectedFile} 
          onClose={closeModal}
          onSuccess={() => { loadData(); closeModal(); }}
        />
      )}
>>>>>>> client
    </div>
  );
}
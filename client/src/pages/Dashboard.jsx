import { useEffect, useState, useCallback } from 'react';
import Sidebar from '../components/layout/Sidebar';
import Navbar from '../components/layout/Navbar';
import FileCard from '../components/FileCard';
import ShareModal from '../components/ShareModal';
import RenameModal from '../components/RenameModal';
import { filesApi } from '../api/files';
import { useUpload } from '../hooks/useUpload';
import { useAuth } from '../context/AuthContext';

export default function Dashboard() {
  const [activeTab, setActiveTab] = useState('owned');
  const [files, setFiles] = useState([]);
  const [shortcutIds, setShortcutIds] = useState(new Set());
  const [publicFileIds, setPublicFileIds] = useState(new Set());
  const [selectedFile, setSelectedFile] = useState(null);
  const [showShareModal, setShowShareModal] = useState(false);
  const [showRenameModal, setShowRenameModal] = useState(false);
  const { uploadFile, progress, isUploading } = useUpload();
  const { logout } = useAuth();

  // Unified data loader based on active tab
  const loadData = useCallback(async () => {
    try {
      let data;
      if (activeTab === 'owned') data = await filesApi.getOwned();
      else if (activeTab === 'shared') data = await filesApi.getShared();
      else if (activeTab === 'shortcuts') data = await filesApi.getShortcuts();
      
      let fileList = data?.files || [];
      
      // Sort by date (newest first) - assuming created_at or updated_at field exists
      fileList.sort((a, b) => {
        const dateA = new Date(a.created_at || a.updated_at || 0);
        const dateB = new Date(b.created_at || b.updated_at || 0);
        return dateB - dateA; // newest first
      });
      
      setFiles(fileList);
      
      // Update shortcut IDs set
      if (activeTab === 'shortcuts') {
        setShortcutIds(new Set(fileList.map(f => f.file_id)));
      } else {
        // Load shortcuts separately to check which files are shortcuts
        try {
          const shortcutsData = await filesApi.getShortcuts();
          const shortcutFileIds = (shortcutsData?.files || []).map(f => f.file_id);
          setShortcutIds(new Set(shortcutFileIds));
        } catch (err) {
          console.error('Failed to load shortcuts:', err);
        }
      }
      
      // TODO: Load public file IDs if backend provides this info
      // For now, check is_public field if it exists in file object
      const publicIds = new Set();
      fileList.forEach(f => {
        if (f.is_public || f.public_url) {
          publicIds.add(f.file_id);
        }
      });
      setPublicFileIds(publicIds);
    } catch (err) {
      console.error("Failed to fetch:", err);
      setFiles([]);
    }
  }, [activeTab]);

  useEffect(() => {
    loadData();
  }, [loadData]);

  // Handle local file selection and trigger upload
  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (file) {
      const success = await uploadFile(file);
      if (success) {
        // Reset input and reload list
        e.target.value = null; 
        loadData();
      }
    }
  };

  // Secure download logic using S3 Presigned URL
  const handleDownload = async (fileId, fileName) => {
    try {
      const data = await filesApi.getDownloadUrl(fileId);
      if (data && data.download_url) {
        const link = document.createElement('a');
        link.href = data.download_url;
        // Setting 'download' attribute helps hint the filename to the browser
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      }
    } catch (err) {
      alert("Download failed: " + err.message);
    }
  };

  const handleDelete = async (fileId) => {
    if (window.confirm("Move this file to trash?")) {
      try {
        await filesApi.delete(fileId);
        loadData();
      } catch (err) {
        alert("Delete failed: " + err.message);
      }
    }
  };

  const handleShare = (file) => {
    setSelectedFile(file);
    setShowShareModal(true);
  };

  const handleRename = (file) => {
    setSelectedFile(file);
    setShowRenameModal(true);
  };

  const handleShortcut = async (fileId) => {
    try {
      const isShortcut = shortcutIds.has(fileId);
      if (isShortcut) {
        await filesApi.removeShortcut(fileId);
      } else {
        await filesApi.addShortcut(fileId);
      }
      loadData(); // Reload to update shortcut status
    } catch (err) {
      alert("Shortcut operation failed: " + err.message);
    }
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
      } else {
        const result = await filesApi.makePublic(file.file_id);
        setPublicFileIds(prev => new Set(prev).add(file.file_id));
        if (result?.public_url) {
          alert(`File is now public!\nPublic URL: ${result.public_url}`);
        }
      }
      loadData();
    } catch (err) {
      alert("Public access operation failed: " + err.message);
    }
  };

  return (
    <div className="flex h-screen bg-[#0d1117] text-gray-200 overflow-hidden">
      {/* 1. Left Sidebar Navigation */}
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      
      {/* 2. Main Area Container */}
      <div className="flex-1 flex flex-col min-w-0">
        
        {/* 3. Global Top Navbar */}
        <Navbar />

        {/* 4. Action Header */}
        <header className="h-14 border-b border-[#30363d] bg-[#0d1117] flex items-center justify-between px-8 shrink-0">
          <h2 className="text-white font-semibold text-lg capitalize">
            {activeTab.replace('-', ' ')}
          </h2>
          
          <div className="flex items-center gap-3">
            <input 
              type="file" 
              id="up-dash" 
              className="hidden" 
              onChange={handleUpload} 
              disabled={isUploading} 
            />
            <label 
              htmlFor="up-dash" 
              className={`px-4 py-1.5 rounded-md text-sm font-bold cursor-pointer transition-all shadow-sm ${
                isUploading 
                  ? 'bg-gray-800 text-gray-500 cursor-not-allowed' 
                  : 'bg-blue-600 hover:bg-blue-500 text-white active:scale-95'
              }`}
            >
              {isUploading ? `Uploading ${progress}%` : 'Upload New File'}
            </label>
          </div>
        </header>

        {/* 5. Scrollable Content Body */}
        <section className="flex-1 overflow-y-auto p-8 custom-scrollbar">
          
          {/* Active Upload Progress Bar */}
          {isUploading && (
            <div className="mb-8 p-4 bg-blue-600/5 border border-blue-500/20 rounded-xl animate-in fade-in slide-in-from-top-2">
              <div className="flex justify-between items-end text-xs mb-2">
                <div className="flex items-center gap-2">
                  <div className="w-2 h-2 bg-blue-500 rounded-full animate-pulse"></div>
                  <span className="text-blue-400 font-semibold uppercase tracking-wider">Chunking & Uploading</span>
                </div>
                <span className="text-blue-400 font-mono">{progress}%</span>
              </div>
              <div className="w-full bg-[#161b22] h-1.5 rounded-full overflow-hidden">
                <div 
                  className="bg-blue-500 h-full transition-all duration-500 ease-out" 
                  style={{ width: `${progress}%` }} 
                />
              </div>
            </div>
          )}

          {/* File Grid Area */}
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-5">
            {files.map(file => (
              <FileCard 
                key={file.file_id} 
                file={file} 
                onDownload={handleDownload}
                onDelete={handleDelete}
                onShare={handleShare}
                onRename={handleRename}
                onShortcut={handleShortcut}
                onTogglePublic={handleTogglePublic}
                isShortcut={shortcutIds.has(file.file_id)}
                isPublic={publicFileIds.has(file.file_id)}
              />
            ))}
          </div>

          {/* Modals */}
          {showShareModal && selectedFile && (
            <ShareModal 
              file={selectedFile} 
              onClose={() => {
                setShowShareModal(false);
                setSelectedFile(null);
                loadData();
              }} 
            />
          )}
          
          {showRenameModal && selectedFile && (
            <RenameModal 
              file={selectedFile} 
              onClose={() => {
                setShowRenameModal(false);
                setSelectedFile(null);
              }}
              onSuccess={loadData}
            />
          )}

          {/* Empty State */}
          {files.length === 0 && !isUploading && (
            <div className="flex flex-col items-center justify-center py-40 opacity-40">
              <div className="text-7xl mb-6">📁</div>
              <p className="text-xl font-medium">Your vault is empty</p>
              <p className="text-sm mt-1">Upload files to get started</p>
            </div>
          )}
        </section>
      </div>
    </div>
  );
}
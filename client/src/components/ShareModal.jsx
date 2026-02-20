import { useState, useEffect } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';
import { filesApi } from '../api/files';

export default function ShareModal({ file, onClose }) {
  const [emailInput, setEmailInput] = useState('');
  const [selectedPermission, setSelectedPermission] = useState('viewer');
  const [pendingRecipients, setPendingRecipients] = useState([]);
  const [existingShares, setExistingShares] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchShares = async () => {
    try {
      setLoading(true);
      const data = await filesApi.getShares(file.file_id);
      setExistingShares(data?.shares || data || []);
    } catch (err) {
      console.error("Error fetching shares:", err);
      setExistingShares([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => { 
    if (file?.file_id) {
      fetchShares(); 
    }
  }, [file?.file_id]);

  const addRecipientToList = () => {
    if (!emailInput.includes('@')) return;
    if (pendingRecipients.find(r => r.email === emailInput)) return;
    setPendingRecipients([...pendingRecipients, { email: emailInput.trim(), permission: selectedPermission }]);
    setEmailInput('');
  };

  const handleBatchShare = async () => {
    try {
      await filesApi.addShares(file.file_id, pendingRecipients);
      setPendingRecipients([]);
      fetchShares();
      alert("Invites sent successfully");
    } catch (err) {
      alert("Share failed: " + err.message);
    }
  };

  const handleUpdatePermission = async (userId, newPermission) => {
    try {
      await filesApi.updateShare(file.file_id, userId, newPermission);
      fetchShares(); // Refresh the list
    } catch (err) {
      alert("Failed to update permission: " + err.message);
    }
  };

  const handleRemove = async (userId) => {
    if (!confirm('Remove access for this user?')) return;
    try {
      await filesApi.removeShare(file.file_id, userId);
      fetchShares();
    } catch (err) {
      alert("Failed to remove share: " + err.message);
    }
  };

  return (
    <div className="fixed inset-0 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4 z-50">
      <div className="bg-[#161b22] border border-[#30363d] w-full max-w-lg rounded-2xl shadow-2xl">
        <div className="p-6 border-b border-[#30363d] flex justify-between items-center">
          <h3 className="text-xl font-bold text-white font-sans">Share "{file.name}"</h3>
          <button onClick={onClose} className="text-gray-400 hover:text-white text-xl">✕</button>
        </div>

        <div className="p-6 space-y-6">
          <div className="flex gap-2">
            <input 
              type="email" placeholder="user@example.com" value={emailInput}
              onChange={(e) => setEmailInput(e.target.value)}
              className="flex-1 bg-gv-dark border border-[#30363d] rounded-lg px-3 py-2 text-sm text-white focus:border-blue-500 outline-none"
            />
            <select 
              value={selectedPermission}
              onChange={(e) => setSelectedPermission(e.target.value)}
              className="bg-gv-dark border border-[#30363d] rounded-lg px-2 text-xs text-white"
            >
              <option value="viewer">Viewer</option>
              <option value="editor">Editor</option>
            </select>
            <button onClick={addRecipientToList} className="bg-blue-600 px-4 py-2 rounded-lg text-sm font-bold text-white">+</button>
          </div>

          {pendingRecipients.length > 0 && (
            <div className="p-3 bg-blue-600/5 border border-blue-500/20 rounded-xl space-y-2">
              <div className="flex flex-wrap gap-2">
                {pendingRecipients.map(r => (
                  <span key={r.email} className="bg-blue-600/20 text-blue-400 text-[10px] px-2 py-1 rounded-full border border-blue-600/30">
                    {r.email} ({r.permission})
                  </span>
                ))}
              </div>
              <button onClick={handleBatchShare} className="w-full bg-blue-600 text-white text-xs font-bold py-2 rounded-lg">Send Batch Invites</button>
            </div>
          )}

          <div className="space-y-3">
            <h4 className="text-[10px] font-bold text-gray-500 uppercase tracking-widest">People with access</h4>
            {loading ? (
              <div className="text-center py-4 text-gray-400 text-sm">Loading shares...</div>
            ) : existingShares.length === 0 ? (
              <div className="text-center py-4 text-gray-500 text-sm">No one has access yet</div>
            ) : (
              <div className="max-h-64 overflow-y-auto space-y-2">
                {existingShares.map(share => {
                  const userId = share.user_id || share.id || share.email;
                  const currentPermission = share.permission || 'viewer';
                  return (
                    <div key={userId} className="flex items-center gap-2 bg-gv-dark p-3 rounded-xl border border-[#30363d]">
                      <span className="text-xs text-gray-300 truncate flex-1">
                        {share.email || share.user_email || share.user_id || userId}
                      </span>
                      <select
                        value={currentPermission}
                        onChange={(e) => handleUpdatePermission(userId, e.target.value)}
                        className="bg-[#161b22] border border-[#30363d] rounded-lg px-2 py-1 text-xs text-white focus:border-blue-500 outline-none"
                      >
                        <option value="viewer">Viewer</option>
                        <option value="editor">Editor</option>
                      </select>
                      <button
                        onClick={() => handleRemove(userId)}
                        className="text-red-500 hover:text-red-400 text-sm px-2 py-1 transition-colors"
                        title="Remove access"
                      >
                        🗑️
                      </button>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
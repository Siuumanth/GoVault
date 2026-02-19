import { useState, useEffect } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';

export default function ShareModal({ file, onClose }) {
  const [emailInput, setEmailInput] = useState('');
  const [selectedPermission, setSelectedPermission] = useState('viewer');
  const [existingShares, setExistingShares] = useState([]);

  // Load existing permissions
  const fetchShares = async () => {
    try {
      const data = await request(ENDPOINTS.SHARING.LIST(file.file_id));
      setExistingShares(data?.shares || data || []);
    } catch (err) { 
      console.error('Failed to fetch shares:', err);
      setExistingShares([]);
    }
  };

  useEffect(() => { 
    if (file?.file_id) {
      fetchShares(); 
    }
  }, [file?.file_id]);

  const addRecipientToList = async () => {
    if (!emailInput.trim()) return;

    const recipientList = [{
      email: emailInput.trim(),
      permission: selectedPermission
    }];

    try {
      await request(ENDPOINTS.SHARING.ADD(file.file_id), {
        method: 'POST',
        body: JSON.stringify({ recipients: recipientList })
      });
      setEmailInput('');
      setSelectedPermission('viewer');
      fetchShares();
    } catch (err) { 
      alert('Failed to share: ' + (err.message || 'Unknown error'));
    }
  };

  const handleRemoveShare = async (userId) => {
    try {
      await request(ENDPOINTS.SHARING.REMOVE(file.file_id, userId), {
        method: 'DELETE',
      });
      fetchShares();
    } catch (err) {
      alert('Failed to remove share: ' + err.message);
    }
  };

  return (
    <div className="fixed inset-0 bg-black/70 backdrop-blur-sm flex items-center justify-center p-4 z-50">
      <div className="bg-[#161b22] border border-[#30363d] w-full max-w-lg rounded-2xl shadow-2xl">
        <div className="p-6 border-b border-[#30363d] flex justify-between items-center">
          <h3 className="text-xl font-bold text-white">Share "{file.name}"</h3>
          <button onClick={onClose} className="text-gray-400 hover:text-white">✕</button>
        </div>

        <div className="p-6 space-y-6">
          {/* Add Section */}
          <div className="space-y-3">
            <label className="block text-sm font-medium text-gray-400 mb-2">Add people (comma separated)</label>
            <div className="flex gap-2">
              <input 
                type="text" 
                value={emailInput} 
                onChange={(e) => setEmailInput(e.target.value)}
                placeholder="user@example.com, dev@govault.com"
                onKeyPress={(e) => e.key === 'Enter' && addRecipientToList()}
                className="flex-1 bg-[#0d1117] border border-[#30363d] rounded-lg px-3 py-2 text-sm text-white outline-none focus:border-blue-500"
              />
              <select 
                value={selectedPermission}
                onChange={(e) => setSelectedPermission(e.target.value)}
                className="bg-[#0d1117] border border-[#30363d] rounded-lg px-2 py-2 text-xs text-white outline-none"
              >
                <option value="viewer">Viewer</option>
                <option value="editor">Editor</option>
              </select>
              <button 
                onClick={addRecipientToList} 
                className="bg-blue-600 px-4 py-2 rounded-lg text-sm font-bold text-white hover:bg-blue-500 transition-colors"
              >
                Add
              </button>
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-400 mb-2">People with access</label>
            <div className="space-y-2 max-h-40 overflow-y-auto">
              {existingShares.length === 0 ? (
                <p className="text-sm text-gray-500 text-center py-4">No one has access yet</p>
              ) : (
                existingShares.map(share => (
                  <div key={share.user_id || share.id || share.email} className="flex justify-between items-center bg-[#0d1117] p-2 rounded-lg border border-[#30363d]">
                    <span className="text-sm text-gray-300 truncate mr-2">
                      {share.email || share.user_id || share.user_email || 'Unknown'}
                    </span>
                    <div className="flex items-center gap-2">
                      <span className="text-[10px] uppercase bg-gray-800 px-2 py-1 rounded text-gray-400">
                        {share.permission || 'viewer'}
                      </span>
                      <button
                        onClick={() => handleRemoveShare(share.user_id || share.id)}
                        className="text-red-400 hover:text-red-300 text-xs transition-colors"
                        title="Remove access"
                      >
                        ✕
                      </button>
                    </div>
                  </div>
                ))
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

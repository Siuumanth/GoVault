import { useState, useEffect } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';

export default function ShareModal({ file, onClose }) {
  const [emailInput, setEmailInput] = useState('');
  const [selectedPermission, setSelectedPermission] = useState('viewer');
  const [pendingRecipients, setPendingRecipients] = useState([]);
  const [existingShares, setExistingShares] = useState([]);

  // Load existing permissions
  const fetchShares = async () => {
    try {
      const data = await request(ENDPOINTS.FILES.SHARES(file.file_id));
      setExistingShares(data || []);
    } catch (err) {
      console.error("Error fetching shares:", err);
    }
  };

  useEffect(() => { fetchShares(); }, [file]);

  const addRecipientToList = () => {
    if (!emailInput.includes('@')) return;
    if (pendingRecipients.find(r => r.email === emailInput)) return;
    
    setPendingRecipients([...pendingRecipients, { email: emailInput, permission: selectedPermission }]);
    setEmailInput('');
  };

  // Inside ShareModal.jsx

const handleBatchShare = async () => {
  if (pendingRecipients.length === 0) return;

  try {
    // We must wrap the array in an object with the key "recipients"
    // to match the Go DTO: type AddFileSharesRequest struct { Recipients []... }
    const payload = {
      recipients: pendingRecipients // pendingRecipients is already [{email, permission}, ...]
    };

    await request(ENDPOINTS.FILES.SHARES(file.file_id), {
      method: 'POST',
      body: JSON.stringify(payload)
    });

    setPendingRecipients([]);
    fetchShares();
    alert("Shares added successfully");
  } catch (err) {
    // If the backend returns "internal error", it's likely a failure 
    // in s.authClient.ResolveEmails (Auth service is down or User doesn't exist)
    alert("Failed to share: " + err.message);
  }
};

  const handleUpdatePermission = async (userId, newPerm) => {
    try {
      await request(ENDPOINTS.FILES.MANAGE_USER(file.file_id, userId), {
        method: 'PATCH',
        body: JSON.stringify({ permission: newPerm })
      });
      fetchShares();
    } catch (err) { alert(err.message); }
  };

  const handleRevoke = async (userId) => {
    try {
      await request(ENDPOINTS.FILES.MANAGE_USER(file.file_id, userId), {
        method: 'DELETE'
      });
      fetchShares();
    } catch (err) { alert(err.message); }
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
            <div className="flex gap-2">
              <input 
                type="email" placeholder="Add email address..." value={emailInput}
                onChange={(e) => setEmailInput(e.target.value)}
                className="flex-1 bg-[#0d1117] border border-[#30363d] rounded-lg px-3 py-2 text-sm text-white focus:border-blue-500 outline-none"
              />
              <select 
                value={selectedPermission}
                onChange={(e) => setSelectedPermission(e.target.value)}
                className="bg-[#0d1117] border border-[#30363d] rounded-lg px-2 py-2 text-xs text-white outline-none"
              >
                <option value="viewer">Viewer</option>
                <option value="editor">Editor</option>
              </select>
              <button onClick={addRecipientToList} className="bg-[#30363d] px-4 py-2 rounded-lg text-sm font-bold text-white">+</button>
            </div>

            {/* Pending List */}
            {pendingRecipients.length > 0 && (
              <div className="flex flex-wrap gap-2">
                {pendingRecipients.map(r => (
                  <span key={r.email} className="bg-blue-600/20 text-blue-400 text-[10px] px-2 py-1 rounded-full border border-blue-600/30 flex items-center gap-2">
                    {r.email} ({r.permission})
                    <button onClick={() => setPendingRecipients(pendingRecipients.filter(p => p.email !== r.email))}>✕</button>
                  </span>
                ))}
                <button onClick={handleBatchShare} className="ml-auto text-xs text-blue-400 font-bold hover:underline">Send Invites</button>
              </div>
            )}
          </div>

          {/* Existing Access List */}
          <div className="space-y-3">
            <h4 className="text-xs font-bold text-gray-500 uppercase tracking-widest">People with access</h4>
            <div className="max-h-48 overflow-y-auto space-y-2 pr-2 custom-scrollbar">
              {existingShares.map(share => (
                <div key={share.user_id} className="flex justify-between items-center bg-[#0d1117] p-3 rounded-xl border border-[#30363d]">
                  <div className="flex flex-col">
                    <span className="text-sm text-gray-200 truncate w-40">{share.user_id}</span>
                    <span className="text-[10px] text-gray-500 font-mono">Added {new Date(share.created_at).toLocaleDateString()}</span>
                  </div>
                  <div className="flex gap-2">
                    <select 
                      value={share.permission}
                      onChange={(e) => handleUpdatePermission(share.user_id, e.target.value)}
                      className="bg-transparent text-xs text-blue-400 outline-none cursor-pointer"
                    >
                      <option value="viewer">Viewer</option>
                      <option value="editor">Editor</option>
                    </select>
                    <button onClick={() => handleRevoke(share.user_id)} className="text-red-500/50 hover:text-red-500 text-xs px-2">Revoke</button>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
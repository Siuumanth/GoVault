import { useState, useEffect } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';

export default function ShareModal({ file, onClose }) {
  const [emailInput, setEmailInput] = useState('');
  const [selectedPermission, setSelectedPermission] = useState('viewer');
  const [pendingRecipients, setPendingRecipients] = useState([]);
  const [existingShares, setExistingShares] = useState([]);

  const fetchShares = async () => {
    try {
      const data = await request(ENDPOINTS.FILES.SHARES(file.file_id));
      setExistingShares(data || []);
    } catch (err) {
      console.error("Error fetching shares:", err);
    }
  };

  useEffect(() => { if (file?.file_id) fetchShares(); }, [file?.file_id]);

  const addRecipientToList = () => {
    if (!emailInput.includes('@')) return;
    if (pendingRecipients.find(r => r.email === emailInput)) return;
    setPendingRecipients([...pendingRecipients, { email: emailInput.trim(), permission: selectedPermission }]);
    setEmailInput('');
  };

  const handleBatchShare = async () => {
    try {
      // WRAPPING IN THE 'recipients' KEY FOR GO DTO
      await request(ENDPOINTS.FILES.SHARES(file.file_id), {
        method: 'POST',
        body: JSON.stringify({ recipients: pendingRecipients })
      });
      setPendingRecipients([]);
      fetchShares();
      alert("Invites sent successfully");
    } catch (err) {
      alert("Share failed: " + err.message);
    }
  };

  const handleRemove = async (userId) => {
    try {
      await request(ENDPOINTS.FILES.MANAGE_USER(file.file_id, userId), { method: 'DELETE' });
      fetchShares();
    } catch (err) { alert(err.message); }
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
              className="flex-1 bg-[#0d1117] border border-[#30363d] rounded-lg px-3 py-2 text-sm text-white focus:border-blue-500 outline-none"
            />
            <select 
              value={selectedPermission}
              onChange={(e) => setSelectedPermission(e.target.value)}
              className="bg-[#0d1117] border border-[#30363d] rounded-lg px-2 text-xs text-white"
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
            <div className="max-h-48 overflow-y-auto space-y-2">
              {existingShares.map(share => (
                <div key={share.user_id} className="flex justify-between items-center bg-[#0d1117] p-3 rounded-xl border border-[#30363d]">
                  <span className="text-xs text-gray-300 truncate w-48">{share.user_id}</span>
                  <button onClick={() => handleRemove(share.user_id)} className="text-red-500 hover:text-red-400 text-xs px-2">Revoke</button>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
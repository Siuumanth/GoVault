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
      const data = await request(ENDPOINTS.SHARING.LIST(file.file_id));
      setShares(data?.shares || data || []);
    } catch (err) { 
      console.error('Failed to fetch shares:', err);
      setShares([]);
    }
  };

  useEffect(() => { 
    if (file?.file_id) {
      fetchShares(); 
    }
  }, [file?.file_id]);

  const handleAddShares = async (e) => {
    e.preventDefault();
    if (!emails.trim()) return;

    const recipientList = emails.split(',').map(email => ({
      email: email.trim(),
      permission: 'viewer'
    })).filter(r => r.email);

    if (recipientList.length === 0) {
      alert('Please enter at least one valid email');
      return;
    }

    try {
      await request(ENDPOINTS.SHARING.ADD(file.file_id), {
        method: 'POST',
        body: JSON.stringify({ recipients: recipientList })
      });
      setEmails('');
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

  const handleRevoke = async (userId) => {
    try {
      await request(ENDPOINTS.FILES.MANAGE_USER(file.file_id, userId), {
        method: 'DELETE'
      });
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
          </form>

          <div>
            <label className="block text-sm font-medium text-gray-400 mb-2">People with access</label>
            <div className="space-y-2 max-h-40 overflow-y-auto">
              {shares.map(share => (
                <div key={share.user_id} className="flex justify-between items-center bg-[#0d1117] p-2 rounded-lg border border-[#30363d]">
                  <span className="text-sm text-gray-300 truncate mr-2">{share.user_id}</span>
                  <span className="text-[10px] uppercase bg-gray-800 px-2 py-1 rounded text-gray-400">{share.permission}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
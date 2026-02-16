import { useState, useEffect } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';

export default function ShareModal({ file, onClose }) {
  const [emails, setEmails] = useState('');
  const [shares, setShares] = useState([]);
  const [isPublic, setIsPublic] = useState(false);

  const fetchShares = async () => {
    try {
      const data = await request(ENDPOINTS.SHARING.LIST(file.file_id));
      setShares(data || []);
    } catch (err) { console.error(err); }
  };

  useEffect(() => { fetchShares(); }, [file]);

  const handleAddShares = async (e) => {
    e.preventDefault();
    const recipientList = emails.split(',').map(email => ({
      email: email.trim(),
      permission: 'viewer'
    }));

    try {
      await request(ENDPOINTS.SHARING.LIST(file.file_id), {
        method: 'POST',
        body: JSON.stringify({ recipients: recipientList })
      });
      setEmails('');
      fetchShares();
    } catch (err) { alert(err.message); }
  };

  return (
    <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
      <div className="bg-[#161b22] border border-[#30363d] w-full max-w-md rounded-2xl overflow-hidden shadow-2xl">
        <div className="p-6 border-b border-[#30363d] flex justify-between items-center">
          <h3 className="text-xl font-bold text-white">Share "{file.name}"</h3>
          <button onClick={onClose} className="text-gray-500 hover:text-white">✕</button>
        </div>

        <div className="p-6 space-y-6">
          <form onSubmit={handleAddShares}>
            <label className="block text-sm font-medium text-gray-400 mb-2">Add people (comma separated)</label>
            <div className="flex gap-2">
              <input 
                type="text" value={emails} onChange={(e) => setEmails(e.target.value)}
                placeholder="user@example.com, dev@govault.com"
                className="flex-1 bg-[#0d1117] border border-[#30363d] rounded-lg px-3 py-2 text-sm text-white outline-none focus:border-blue-500"
              />
              <button className="bg-blue-600 px-4 py-2 rounded-lg text-sm font-bold text-white">Add</button>
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
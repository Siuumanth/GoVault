import { useAuth } from './context/AuthContext';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import FilePreview from './pages/FilePreview';

function App() {
  const { user } = useAuth();
  const path = typeof window !== 'undefined' ? window.location.pathname : '';
  
  // Handle public file preview route
  if (path.startsWith('/f/')) {
    return <FilePreview />;
  }
  
  return user ? <Dashboard /> : <Login />;
}

export default App;
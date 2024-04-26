import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import NavigationBar from './layouts/NavigationBar';

import RootLayout from './layouts/RootLayout';
import Dashboard from './pages/Dashboard';
import Profile from './pages/Profile';
import BFS from './pages/BFS';
import IDS from './pages/IDS';
import HowToUse from './pages/HowToUse';

function App() {
  return (
    <Router>
      <NavigationBar />
      <Routes>
        <Route path="/" element={<RootLayout />}>
          <Route index element={<Dashboard />} />
          <Route path="profile" element={<Profile />} />
          <Route path="bfs-page" element={<BFS />} />
          <Route path="ids-page" element={<IDS />} />
          <Route path="how-to-use" element={<HowToUse />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;

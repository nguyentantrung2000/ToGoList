import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import './App.css';
import HomePage from './page/HomePage';
import MessagePage from "./page/MessagePage";
import SideBar from './components/SideBar';
import { Layout, Flex } from 'antd';
import Header from "./components/Header";
const { Content } = Layout;


function App() {
  const headerStyle = {
    textAlign: 'center',
    color: '#fff',
    height: 80,
    paddingInline: 48,
    lineHeight: '64px',
    backgroundColor: '#4096ff',
  };
  const layoutStyle = {
    borderRadius: 8,
    overflow: 'hidden',
  };

  return (
    <Router>
      <div>
        <Layout style={layoutStyle}>
          <SideBar />
          <Layout>
            <Header/>
            <Content>
              <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="message" element={<MessagePage />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </div>
    </Router>
  );
}

export default App;

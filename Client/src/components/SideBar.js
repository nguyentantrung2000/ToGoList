import React from 'react';
import { AppstoreOutlined, MailOutlined, SettingOutlined, HomeOutlined } from '@ant-design/icons';
import './config/style.css';
import { Layout, Menu } from 'antd';
import { Link,  } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const { Header, Footer, Sider, Content } = Layout;


const items = [
    { icon: <HomeOutlined />, path: "/" },
    { icon: <MailOutlined />, path: "/message" },
    { icon: <SettingOutlined />, path: "/settings" },
];

const SideBar = () => {
    const navigate = useNavigate();

    const onClick = (e) => {
        const item = items[e.key];
        if (item) {
            navigate(item.path);
        }
    };
    return (
        <Sider
            style={{
                height: '100vh',
            }}
        >
            <Menu
                className='menu'
                onClick={onClick}
                style={{
                    backgroundColor: '#1C1D22',
                    height: '100vh',
                    borderRight: 0,
                }}
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                mode="inline"
                align="center"
            >
                {items.map((item, index) => (
                    <Menu.Item key={index}>
                        <Link to={item.path}>
                            {item.icon}
                        </Link>
                    </Menu.Item>
                ))}
            </Menu>
        </Sider>
    );
};
export default SideBar;
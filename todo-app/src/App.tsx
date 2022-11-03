import { Routes, Route, Link } from "react-router-dom";
import { useAuthorizer } from "@authorizerdev/authorizer-react";
import Home from "./pages/home";
import Dashboard from "./pages/dashboard";
import Login from "./pages/login";
import { useNavigate } from "react-router-dom";

export default function App() {
  const { user, loading, authorizerRef, setUser, setToken } = useAuthorizer();
  const navigate = useNavigate();

  const logout = async () => {
    await authorizerRef.logout();
    setUser(null);
    setToken(null);
    navigate("/");
  };

  if (loading) {
    return <h1> Loading...</h1>;
  }

  return (
    <div>
      {user ? (
        <div>
          <nav className="navbar">
            <div className="logo-container">
              <img src="aws-logo.png" alt="logo" className="logo" />
              <h1>Task Manager</h1>
            </div>
            <div className="user-info">
              <span>👋 {user.email}&nbsp; | &nbsp; </span>
              <button onClick={logout} className="link-btn">
                Logout
              </button>
            </div>
          </nav>
          <div className="container">
            <Routes>
              <Route path="/" element={<Dashboard />} />
            </Routes>
          </div>
        </div>
      ) : (
        <div>
          <nav className="navbar">
            <div className="logo-container">
              <img src="aws-logo.png" alt="logo" className="logo" />
              <h1>Task Manager</h1>
            </div>
            <div>
              <Link to="/">Home</Link> | <Link to="/login">Login</Link>
            </div>
          </nav>
          <br />
          <div className="container">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />
            </Routes>
          </div>
        </div>
      )}
    </div>
  );
}

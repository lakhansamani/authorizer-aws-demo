import { Link } from "react-router-dom";

export default function Home() {
  return (
    <>
      <h1>Welcome to AWS Community Event</h1>
      <h3>Login And Start Managing Your Tasks</h3>
      <Link to="/login" style={{ fontSize: 32 }}>
        Login
      </Link>
    </>
  );
}

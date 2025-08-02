import { BrowserRouter, Routes, Route } from "react-router-dom";
import JobSubmitter from "./JobSubmitter";
import JobList from "./JobList";
import "./App.css";

export default function App() {
  return (
    <BrowserRouter>
      <div
        style={{ width: "100%", maxWidth: 600, margin: "auto", padding: 20 }}
      >
        <h1>Distributed Task Runner UI</h1>
        <Routes>
          <Route path="/" element={<JobSubmitter />} />
          <Route path="/jobs" element={<JobList />} />
        </Routes>
        <footer style={{ textAlign: "center", marginTop: 50 }}>
          <p>© 2025 Herald Inyang</p>
        </footer>
      </div>
    </BrowserRouter>
  );
}

import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import "./JobList.css";

type Job = {
  id: string;
  payload: string;
  status: string;
  created_at: string;
  updated_at: string;
};

export default function JobList() {
  const [jobs, setJobs] = useState<Job[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        const res = await fetch("/jobs");
        if (!res.ok) throw new Error("Failed to fetch jobs");
        const data = await res.json();
        setJobs(data.reverse()); // most recent first
      } catch (err: any) {
        setError(err.message || "Unknown error");
      } finally {
        setLoading(false);
      }
    };

    fetchJobs();
  }, []);

  if (loading) return <p>Loading jobs...</p>;
  if (error) return <p style={{ color: "red" }}>{error}</p>;

  const clearHistory = async () => {
    if (!window.confirm("Are you sure you want to clear all job history?"))
      return;

    try {
      const res = await fetch("/jobs", { method: "DELETE" });
      if (res.ok) {
        setJobs([]); // clear local job list
      } else {
        alert("Failed to clear job history");
      }
    } catch {
      alert("Failed to clear job history");
    }
  };

  return (
    <>
      <Link to="/">
        <button style={{ marginBottom: 20 }}>← Back</button>
      </Link>
      <button
        onClick={clearHistory}
        className="clear-history-btn"
        style={{ marginBottom: 20 }}
      >
        Clear History
      </button>
      <div style={{ maxWidth: 600, margin: "40px auto", padding: 20 }}>
        <h2 style={{ textAlign: "center" }}>Job History</h2>
        {jobs.length === 0 ? (
          <p style={{ textAlign: "center" }}>No jobs found.</p>
        ) : (
          <ul style={{ listStyle: "none", padding: 0 }}>
            {jobs.map((job) => (
              <li
                key={job.id}
                style={{
                  padding: 12,
                  marginBottom: 12,
                  borderRadius: 6,
                  background: "#f9f9f9",
                  borderLeft: `4px solid ${
                    job.status === "completed"
                      ? "green"
                      : job.status === "failed"
                      ? "red"
                      : job.status === "running"
                      ? "orange"
                      : "#aaa"
                  }`,
                }}
              >
                <div>
                  <strong>Command:</strong> {job.payload}
                </div>
                <div>
                  <strong>Status:</strong> {job.status}
                </div>
                <div>
                  <small>{new Date(job.created_at).toLocaleString()}</small>
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>
    </>
  );
}

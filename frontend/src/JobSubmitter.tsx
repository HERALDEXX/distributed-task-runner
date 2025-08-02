import { useState } from "react";
import { Link } from "react-router-dom";
import "./JobSubmitter.css";

type JobResponse = {
  id: string;
  status: string;
};

export default function JobSubmitter() {
  const [payload, setPayload] = useState("");
  const [response, setResponse] = useState<JobResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const submitJob = async () => {
    setLoading(true);
    setError(null);
    setResponse(null);

    try {
      const res = await fetch("/jobs", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ payload }),
      });
      if (!res.ok) throw new Error("Failed to enqueue job");

      const data: JobResponse = await res.json();
      setResponse(data);
    } catch (err: any) {
      setError(err.message || "Unknown error");
    } finally {
      setLoading(false);
    }
  };
  return (
    <>
      <div className="container">
        <h2 style={{ textAlign: "center" }}>Submit a Job</h2>
        <textarea
          rows={4}
          placeholder="Enter command or script to run"
          value={payload}
          onChange={(e) => setPayload(e.target.value)}
        />
        <button onClick={submitJob} disabled={loading || !payload.trim()}>
          {loading ? "Submitting..." : "Submit"}
        </button>
        {response && (
          <div className="job-output">
            <strong>Job ID:</strong> {response.id}
            <br />
            <strong>Status:</strong> {response.status}
          </div>
        )}
        {error && <div className="error">{error}</div>}
      </div>
      <Link to="/jobs" className="view-history-link">
        <button>View Job History</button>
      </Link>
    </>
  );
}

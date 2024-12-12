import React, { useEffect, useState } from "react";
import CodeEditor from "./components/CodeEditor";
import Output from "./components/Output";
import { submitCode } from "./services/api";
import { v4 as uuid } from 'uuid'

function App() {
  const [output, setOutput] = useState(null);
  const [submissionId, setSubmissionId] = useState("123");

  const handleCodeSubmit = async (code, language) => {
    const submission = await submitCode(code, language,submissionId);
    // setSubmissionId(submission.submission_id);
  };

  useEffect(()=>{
    setSubmissionId(uuid())
  },[])

  return (
    <div style={{ padding: "20px" }}>
      <h1>Online Code Compiler</h1>
      <CodeEditor onSubmit={handleCodeSubmit} />
      <Output submissionId={submissionId} onResult={setOutput} />
      {output && (
        <div>
          <h3>Execution Output</h3>
          <pre>{output.output}</pre>
          {output.error && <pre style={{ color: "red" }}>{output.error}</pre>}
        </div>
      )}
    </div>
  );
}

export default App;

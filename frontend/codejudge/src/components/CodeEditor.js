import React, { useState } from "react";
import { Editor } from "@monaco-editor/react";

function CodeEditor({ onSubmit,submissionId }) {
  const [code, setCode] = useState("// Write your code here");
  const [language, setLanguage] = useState("javascript");

  const handleEditorChange = (value) => {
    setCode(value); // Update the code state when the editor content changes
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(code, language,submissionId);
  };

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: "20px" }}>
      <div style={{ height: "70vh", border: "1px solid #ddd", marginBottom: "10px" }}>
        <Editor
          height="100%"
          defaultLanguage={language}
          defaultValue={code}
          theme="vs-dark"
          onChange={handleEditorChange}
          language={language}
        />
      </div>
      <div>
        <select
          value={language}
          onChange={(e) => setLanguage(e.target.value)}
          style={{ marginRight: "10px" }}
        >
          <option value="javascript">JavaScript</option>
          <option value="python">Python</option>
          <option value="go">Go</option>
        </select>
        <button type="submit">Run Code</button>
      </div>
    </form>
  );
}

export default CodeEditor;

import React, { useEffect } from "react";
import { connectWebSocket } from "../services/websocket";

function Output({ submissionId, onResult }) {
  useEffect(() => {
    if (!submissionId) {
      console.log("No submissionId available yet.");
      return;
    }

    console.log(`Establishing WebSocket connection for submissionId: ${submissionId}`);
    const ws = connectWebSocket(submissionId);

    ws.onopen = () => {
      console.log("WebSocket connection established.");
    };

    ws.onmessage = (event) => {
      console.log("WebSocket message received:", event.data);
      const data = JSON.parse(event.data);
      onResult(data);
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    // Close the WebSocket connection when the component unmounts
    return () => {
      console.log("Cleaning up WebSocket connection.");
      ws.close();
    };
  }, [submissionId, onResult]);

  return null; // This component doesn't render anything directly
}

export default Output;

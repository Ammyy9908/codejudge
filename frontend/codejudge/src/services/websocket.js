export function connectWebSocket(submissionId) {
    const ws = new WebSocket(`ws://localhost:8000/ws?submission_id=${submissionId}`);
    return ws;
  }
  
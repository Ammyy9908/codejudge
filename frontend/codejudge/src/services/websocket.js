export function connectWebSocket(submissionId) {
    const ws = new WebSocket(`ws://localhost:8081/ws?submission_id=${submissionId}`);
    return ws;
  }
  
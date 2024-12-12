export async function submitCode(code, language,id) {
    console.log("Submitting code:", code, language);
    const response = await fetch("http://localhost:8080/submit", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ code, language,id }),
    });
  
    if (!response.ok) {
      throw new Error("Failed to submit code");
    }
  
    return response.json();
  }
  
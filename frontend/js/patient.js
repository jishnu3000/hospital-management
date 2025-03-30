document.addEventListener("DOMContentLoaded", async function () {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("Unauthorized! Please login.");
    window.location.href = "login.html";
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/api/patients", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Unauthorized or invalid token");
    }

    const patients = await response.json();
    document.getElementById("patients-list").innerHTML = patients
      .map(
        (patient) => `
            <p>${patient.name} - ${patient.age} years old</p>
        `
      )
      .join("");
  } catch (error) {
    console.error("Error fetching patients:", error.message);
  }
});

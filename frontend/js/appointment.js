document.addEventListener("DOMContentLoaded", async function () {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("Unauthorized! Please login.");
    window.location.href = "login.html";
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/api/appointments", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Unauthorized or invalid token");
    }

    const appointments = await response.json();
    document.getElementById("appointments-list").innerHTML = appointments
      .map(
        (appointment) => `
            <p>Patient: ${appointment.patientName} | Doctor: ${appointment.doctorName} | Date: ${appointment.date}</p>
        `
      )
      .join("");
  } catch (error) {
    console.error("Error fetching appointments:", error.message);
  }
});

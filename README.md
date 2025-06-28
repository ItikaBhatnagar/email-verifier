# Email Verifier

## What is Email Verifier?

The **Email Verifier** is a simple web-based tool that checks whether a domain has valid email configurations. Specifically, it checks:

-  **MX (Mail Exchange)** — Is the domain capable of receiving emails?  
-  **SPF (Sender Policy Framework)** — Does the domain have rules to stop fake senders?  
-  **DMARC (Domain-based Message Authentication, Reporting & Conformance)** — Does the domain have email protection policies against spam and phishing?

You enter a domain (like `gmail.com`), and it tells you if the domain has these email-related records configured correctly.

---

## Features

- Check **MX** records (can receive emails)  
- Check **SPF** records (helps stop fake senders)  
- Check **DMARC** records (spam and phishing protection)  
- Simple web interface with instant results  


---

## Project Structure
email-verifier/
├── main.go → Backend Go server (API + static file server)
├── static/
│ ├── index.html → Frontend web interface
│ └── style.css → (Optional) Styling for the frontend
├── README.md → Documentation

### Steps to Run Locally:

1. **Clone the repository:**
```bash
git clone https://github.com/yourusername/email-verifier.git
cd email-verifier
go run main.go
** Open Breowser and Go to 
http://localhost:8080

![Alt text](images/demo.png)


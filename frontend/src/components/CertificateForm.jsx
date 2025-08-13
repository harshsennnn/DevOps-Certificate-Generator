// // import React, { useState, useEffect } from "react";
// // import { fetchTemplates, generateCertificate } from "../api";

// // export default function CertificateForm({ onPreview }) {
// //   const [templates, setTemplates] = useState([]);
// //   const [formData, setFormData] = useState({
// //     name: "",
// //     description: "",
// //     template_id: ""
// //   });

// //   useEffect(() => {
// //     fetchTemplates().then(data => setTemplates(data));
// //   }, []);

// //   const handleChange = (e) => {
// //     setFormData(prev => ({ ...prev, [e.target.name]: e.target.value }));
// //     if (onPreview) onPreview({ ...formData, [e.target.name]: e.target.value });
// //   };

// //   const handleDownload = async () => {
// //     const res = await generateCertificate(formData);
// //     const blob = new Blob([res.data], { type: "application/pdf" });
// //     const link = document.createElement("a");
// //     link.href = URL.createObjectURL(blob);
// //     link.download = "certificate.pdf";
// //     link.click();
// //   };

// //   const handleReset = () => {
// //     setFormData({ name: "", description: "", template_id: "" });
// //     if (onPreview) onPreview(null);
// //   };

// //   return (
// //     <div className="form-container">
// //       <input
// //         type="text"
// //         name="name"
// //         placeholder="Enter Name"
// //         value={formData.name}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="text"
// //         name="description"
// //         placeholder="Enter Description"
// //         value={formData.description}
// //         onChange={handleChange}
// //       />

// //       <select
// //         name="template_id"
// //         value={formData.template_id}
// //         onChange={handleChange}
// //       >
// //         <option value="">Select Template</option>
// //         {templates.map(t => (
// //           <option key={t.template_id} value={t.template_id}>{t.name}</option>
// //         ))}
// //       </select>

// //       <div className="button-row">
// //         <button onClick={handleDownload} disabled={!formData.template_id}>Download</button>
// //         <button onClick={handleReset}>Reset</button>
// //       </div>
// //     </div>
// //   );
// // }


// // import React, { useState, useEffect } from "react";
// // import { fetchTemplates, generateCertificate } from "../api";

// // export default function CertificateForm() {
// //   const [templates, setTemplates] = useState([]);
// //   const [formData, setFormData] = useState({
// //     name: "",
// //     company: "",
// //     position: "",
// //     duration: "",
// //     start_date: "",
// //     end_date: "",
// //     download_date: "",
// //     template_id: "",
// //     template: null
// //   });

// //   // Fetch templates
// //   useEffect(() => {
// //     const loadTemplates = async () => {
// //       try {
// //         const data = await fetchTemplates();
// //         setTemplates(data);
// //       } catch (error) {
// //         console.error("Error fetching templates:", error);
// //       }
// //     };
// //     loadTemplates();
// //   }, []);

// //   // Handle input changes
// //   const handleChange = (e) => {
// //     const { name, value } = e.target;
// //     let newFormData = { ...formData, [name]: value };

// //     // // If selecting template, attach full template object
// //     // if (name === "template_id") {
// //     //   const selectedTemplate = templates.find(t => t.template_id === value);
// //     //   newFormData.template = selectedTemplate || null;
// //     // }

// //     setFormData(newFormData);
// //   };

// //   // Send to Go API and download PDF
// //   const handleDownload = async () => {
// //     try {
// //       const res = await generateCertificate(formData);
// //       const blob = new Blob([res.data], { type: "application/pdf" });
// //       const link = document.createElement("a");
// //       link.href = URL.createObjectURL(blob);
// //       link.download = "certificate.pdf";
// //       link.click();
// //     } catch (error) {
// //       console.error("Error generating certificate:", error);
// //     }
// //   };

// //   // Reset form
// //   const handleReset = () => {
// //     setFormData({
// //       name: "",
// //       company: "companyA",
// //       position: "",
// //       duration: "",
// //       start_date: "",
// //       end_date: "",
// //       download_date: "",
// //       template_id: "",
// //       template: null
// //     });
// //   };

// //   return (
// //     <div className="form-container">
// //       <input
// //         type="text"
// //         name="name"
// //         placeholder="Enter Name"
// //         value={formData.name}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="text"
// //         name="company"
// //         placeholder="Enter Company"
// //         value={formData.company}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="text"
// //         name="position"
// //         placeholder="Enter Position"
// //         value={formData.position}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="text"
// //         name="duration"
// //         placeholder="Enter Duration"
// //         value={formData.duration}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="date"
// //         name="start_date"
// //         placeholder="Start Date"
// //         value={formData.start_date}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="date"
// //         name="end_date"
// //         placeholder="End Date"
// //         value={formData.end_date}
// //         onChange={handleChange}
// //       />
// //       <input
// //         type="date"
// //         name="download_date"
// //         placeholder="Download Date"
// //         value={formData.download_date}
// //         onChange={handleChange}
// //       />

// //       <select
// //         name="template_id"
// //         value={formData.template_id}
// //         onChange={handleChange}
// //       >
// //         <option value="">Select Template</option>
// //         {templates.map(t => (
// //           <option key={t.template_id} value={t.template_id}>
// //             {t.name}
// //           </option>
// //         ))}
// //       </select>

// //       <div className="button-row">
// //         <button onClick={handleDownload} disabled={!formData.template_id}>
// //           Download
// //         </button>
// //         <button onClick={handleReset}>Reset</button>
// //       </div>
// //     </div>
// //   );
// // }

// // src/CertificateForm.jsx
// import React, { useState } from "react";

// export default function CertificateForm() {
//   const [formData, setFormData] = useState({
//     name: "",
//     company: "companyA",
//     position: "",
//     duration: "",
//     start_date: "",
//     end_date: "",
//     download_date: "",
//   });

//   // Update form fields
//   const handleChange = (e) => {
//     setFormData({
//       ...formData,
//       [e.target.name]: e.target.value,
//     });
//   };

//   // Handle form submission
//   const handleDownload = async (e) => {
//     e.preventDefault();

//     try {
//       const res = await fetch("http://localhost:8081/generate-certificate", {
//         method: "POST",
//         headers: { "Content-Type": "application/json" },
//         body: JSON.stringify(formData),
//       });

//       if (!res.ok) {
//         throw new Error(`Error: ${res.status}`);
//       }

//       const blob = await res.blob();
//       const url = window.URL.createObjectURL(blob);

//       const a = document.createElement("a");
//       a.href = url;
//       a.download = "certificate.pdf";
//       document.body.appendChild(a);
//       a.click();
//       a.remove();

//       window.URL.revokeObjectURL(url);
//     } catch (error) {
//       console.error("Error generating certificate:", error);
//     }
//   };

//   return (
//     <form onSubmit={handleDownload} style={{ maxWidth: "400px", margin: "auto" }}>
//       <h2>Generate Certificate</h2>

//       <label>Name:</label>
//       <input name="name" value={formData.name} onChange={handleChange} required />

//       <label>Company:</label>
//       <select name="company" value={formData.company} onChange={handleChange}>
//         <option value="companyA">Company A</option>
//         <option value="companyB">Company B</option>
//         <option value="companyC">Company C</option>
//       </select>

//       <label>Position:</label>
//       <input name="position" value={formData.position} onChange={handleChange} required />

//       <label>Duration:</label>
//       <input name="duration" value={formData.duration} onChange={handleChange} required />

//       <label>Start Date:</label>
//       <input type="date" name="start_date" value={formData.start_date} onChange={handleChange} required />

//       <label>End Date:</label>
//       <input type="date" name="end_date" value={formData.end_date} onChange={handleChange} required />

//       <label>Download Date:</label>
//       <input type="date" name="download_date" value={formData.download_date} onChange={handleChange} required />

//       <br />
//       <button type="submit">Download PDF</button>
//     </form>
//   );
// }

import React, { useState } from "react";
// import "App.css"; // <-- Import CSS

export default function CertificateForm() {
  const [formData, setFormData] = useState({
    name: "",
    company: "companyA",
    position: "",
    duration: "",
    start_date: "",
    end_date: "",
    download_date: "",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleDownload = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8081/generate-certificate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status}`);
      }

      const blob = await res.blob();
      const url = window.URL.createObjectURL(blob);

      const a = document.createElement("a");
      a.href = url;
      a.download = "certificate.pdf";
      document.body.appendChild(a);
      a.click();
      a.remove();

      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Error generating certificate:", error);
    }
  };

  return (
    <form className="certificate-form" onSubmit={handleDownload}>
      <h2>Generate Certificate</h2>

      <label>Name:</label>
      <input name="name" value={formData.name} onChange={handleChange} required />

      <label>Company:</label>
      <select name="company" value={formData.company} onChange={handleChange}>
        <option value="companyA">Company A</option>
        <option value="companyB">Company B</option>
        <option value="companyC">Company C</option>
      </select>

      <label>Position:</label>
      <input name="position" value={formData.position} onChange={handleChange} required />

      <label>Duration:</label>
      <input name="duration" value={formData.duration} onChange={handleChange} required />

      <label>Start Date:</label>
      <input type="date" name="start_date" value={formData.start_date} onChange={handleChange} required />

      <label>End Date:</label>
      <input type="date" name="end_date" value={formData.end_date} onChange={handleChange} required />

      <label>Download Date:</label>
      <input type="date" name="download_date" value={formData.download_date} onChange={handleChange} required />

      <div className="button-row">
        <button type="submit">Download PDF</button>
      </div>
    </form>
  );
}

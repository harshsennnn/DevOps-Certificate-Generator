// import { useEffect, useState } from "react";

// export default function CertificatePreview({ selectedCompany, formData, setFormData }) {
//   const [template, setTemplate] = useState(null);

//   useEffect(() => {
//     async function loadTemplate() {
//       try {
//         console.log('data is fetching');
//         const res = await fetch('http://localhost:8081/templates');
//         const data = await res.json();
//         console.log("Template loaded:", data);
//         setTemplate(data);

//         // ✅ Merge template into formData
//         setFormData(prev => ({
//           ...prev,
//           template: data
//         }));

//       } catch (error) {
//         console.error("Failed to load templates:", error);
//       }
//     }
//     loadTemplate();
//   }, [selectedCompany, setFormData]);

//   if (!template) return <p>Loading preview...</p>;

//   return (
//     <div style={{ position: "relative", width: "800px", height: "600px" }}>
//       <img
//         src={`http://localhost:8081/templates/${template.background_image}`}
//         alt="Certificate Background"
//         style={{ position: "absolute", width: "100%", height: "100%" }}
//       />

//       {/* Name field preview */}
//       <div
//         style={{
//           position: "absolute",
//           top: `${template.fields.name.y}px`,
//           left: `${template.fields.name.x || 0}px`,
//           fontSize: template.fields.name.font_size,
//           fontFamily: template.fields.name.font_family,
//           fontWeight: template.fields.name.font_style.includes("B") ? "bold" : "normal",
//           fontStyle: template.fields.name.font_style.includes("I") ? "italic" : "normal",
//           color: `rgb(${template.fields.name.color.join(",")})`
//         }}
//       >
//         {formData?.name}
//       </div>
//     </div>
//   );
// }


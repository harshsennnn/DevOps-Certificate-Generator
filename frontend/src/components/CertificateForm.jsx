import React from "react";

export default function CertificateForm({ formData, setFormData, templates, onTemplateSelect, onSubmit, onReset }) {
    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    return (
        <form onSubmit={(e) => e.preventDefault()} className="form-container">
            <input name="name" placeholder="Name" value={formData.name} onChange={handleChange} required />
            <input name="company" placeholder="Company" value={formData.company} onChange={handleChange} required />
            <input name="position" placeholder="Position" value={formData.position} onChange={handleChange} />
            <input name="duration" placeholder="Duration" value={formData.duration} onChange={handleChange} />
            <input name="start_date" placeholder="Start Date (DD/MM/YYYY)" value={formData.start_date} onChange={handleChange} />
            <input name="end_date" placeholder="End Date (DD/MM/YYYY)" value={formData.end_date} onChange={handleChange} />

            <select value={formData.template} onChange={(e) => onTemplateSelect(e.target.value)}>
                <option value="">Select Template</option>
                {templates.map((tpl) => (
                    <option key={tpl.id} value={tpl.id}>{tpl.name}</option>
                ))}
            </select>

            <div className="btn-row">
                <button type="button" onClick={onSubmit}>Download PDF</button>
                <button type="button" onClick={onReset}>Reset</button>
            </div>
        </form>
    );
}

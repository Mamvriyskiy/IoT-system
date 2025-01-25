import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { codeUser } from "../../api/authApi";
import styles from "./CodeForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const CodeForm: React.FC = () => {
    const [email, setEmail] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            await codeUser({ email });
            alert("Код отправлен на почту!");
            navigate("/login");
        } catch (error) {
            alert("Ошибка отправки кода");
        }
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={globStyles.formHeader}>
                <h1>Восстановление пароля</h1>
            </div>
            <form className={globStyles.registrationForm} onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Введите вашу почту"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <button type="submit">Отправить код</button>
            </form>
        </div>
    );
};

export default CodeForm;

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { verificationUser } from "../../api/authApi";
import styles from "./VerificationForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const VerificationForm: React.FC = () => {
    const [code, setCode] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            await verificationUser({ code });
            alert("Регистрация успешна!");
            navigate("/login");
        } catch (error) {
            alert("Ошибка регистрации");
        }
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={styles.formHeader}>
                <h1>Проверка кода</h1>
            </div>
            <form className={styles.registrationForm} onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Введите код"
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    required
                />
                <button type="submit">Отправить код</button>
            </form>
        </div>
    );
};

export default VerificationForm;

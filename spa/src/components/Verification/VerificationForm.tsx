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
            const urlParams = new URLSearchParams(window.location.search);
            if (urlParams.has('token')) {
                const token = urlParams.get('token') as string; // Утверждение, что token не равен null
                console.log('JWT Token:', token);
                localStorage.setItem('jwtToken', token);
                
                await verificationUser({ code, token }); // token гарантированно string
                alert("Верификация пройдена!");
                navigate("/auth/password");
            } else {
                alert('Нет доступа');
                navigate('/auth/code');
            }
        } catch (error) {
            alert("Ошибка верификации");
            navigate("/auth/code");
        }
        
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={globStyles.formHeader}>
                <h1>Проверка кода</h1>
            </div>
            <form className={globStyles.registrationForm} onSubmit={handleSubmit}>
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

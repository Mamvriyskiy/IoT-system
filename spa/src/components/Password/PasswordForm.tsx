import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { passwordUser } from "../../api/authApi";
import styles from "./PasswordForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const PasswordForm: React.FC = () => {
    const [newPassword, setNewPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        if (newPassword !== repeatPassword) {
            alert("Пароли не совпадают!");
            return;
        }

        try {
            const token = localStorage.getItem('jwtToken');
            if (token) {
                console.log('JWT Token:', token);
                localStorage.setItem('jwtToken', token);
                
                console.log(newPassword, token);
                await passwordUser({ newPassword, token });
                alert("Пароль успешно сменен!");
                navigate("/auth/sign-in");
            } else {
                alert("Ошибка смены пароля");
            }
        } catch (error) {
            alert("Ошибка смены пароля");
        }
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={globStyles.formHeader}>
                <h1>Смена пароля</h1>
            </div>
            <form className={globStyles.registrationForm} onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Введите новый пароль"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                    required
                />
                <input
                    type="text"
                    placeholder="Повторите пароль"
                    value={repeatPassword}
                    onChange={(e) => setRepeatPassword(e.target.value)}
                    required
                />
                <button type="submit">Сменить пароль</button>
            </form>
        </div>
    );
};

export default PasswordForm;

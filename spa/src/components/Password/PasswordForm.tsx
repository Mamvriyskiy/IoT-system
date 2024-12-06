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
            await passwordUser({ newPassword, repeatPassword });
            alert("Регистрация успешна!");
            navigate("/login");
        } catch (error) {
            alert("Ошибка регистрации");
        }
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={styles.formHeader}>
                <h2>Смена пароля</h2>
            </div>
            <form className={styles.registrationForm} onSubmit={handleSubmit}>
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

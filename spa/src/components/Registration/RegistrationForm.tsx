import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../../api/authApi";
import styles from "./RegistrationForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const RegistrationForm: React.FC = () => {
    const [login, setLogin] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            await registerUser({ login, email, password });
            alert("Регистрация успешна!");
            navigate("/login");
        } catch (error) {
            alert("Ошибка регистрации");
        }
    };

    return (
        <div className={globStyles.authContainer}> {}
            <div className={styles.formHeader}>
                <h1>Регистрация</h1>
                <p>Создайте учетную запись для управления вашим умным домом</p>
            </div>
            <form className={styles.registrationForm} onSubmit={handleSubmit}>
                <label>Логин</label>
                <input
                    type="text"
                    placeholder="Логин"
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    required
                />
                <label>Почта</label>
                <input
                    type="email"
                    placeholder="Почта"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <label>Пароль</label>
                <input
                    type="password"
                    placeholder="Пароль"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Зарегистрироваться</button>
            </form>
            <div className={styles.formFooter}>
            <p>Уже есть аккаунт? <a href="/auth/sign-in">Войти</a></p>
        </div>
        </div>
    );
};

export default RegistrationForm;

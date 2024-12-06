import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { loginUser } from "../../api/authApi";
import styles from "./AuthForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const AuthForm: React.FC = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const response = await loginUser({ email, password });
            localStorage.setItem("jwt", response.token);
            navigate("/menu");
        } catch (error) {
            alert("Ошибка авторизации");
        }
    };

    return (
        <div className={globStyles.authContainer}>
            <div className="formHeader">
                <h1>Добро пожаловать</h1>
                <p>Войдите, чтобы управлять вашим умным домом</p>
            </div>
            <form className={styles.loginForm} onSubmit={handleSubmit}>
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
                <button type="submit">Войти</button>
                <div className="formFooter">
                    <p>Нет аккаунта? <a href="auth/sign-up">Зарегистрироваться</a></p>
                    <p><a href="/forgot-password">Забыли пароль?</a></p>
                </div>
            </form>
        </div>
    );
};

export default AuthForm;

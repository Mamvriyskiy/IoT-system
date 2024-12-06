import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { passwordUser } from "../../api/authApi";
import styles from "./DeviceForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const PasswordForm: React.FC = () => {
    const [newPassword, setNewPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        try {
            //await passwordUser({ newPassword, repeatPassword });
            alert("Регистрация успешна!");
            navigate("/login");
        } catch (error) {
            alert("Ошибка регистрации");
        }
    };

    return (
        <div>
            <h1>Устройство:</h1>
            <h2>Характеристики устройства</h2>
            <table>
                <thead>
                    <tr>
                        <th>Характеристика</th>
                        <th>Показатели</th>
                        <th>Единицы измерения</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Вес</td>
                        <td>5</td>
                        <td>Кг</td>
                    </tr>
                </tbody>
            </table>

            <h2>История работы устройств</h2>
            <table>
                <thead>
                    <tr>
                        <th>Дата работы</th>
                        <th>Время работы</th>
                        <th>Потраченные ресурсы</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>02.01.2003</td>
                        <td>29 мин</td>
                        <td>
                            20вт
                        </td>
                    </tr>
                </tbody>
            </table>
            <div style={{ textAlign: "center" }}>
                <button className={styles.addButton}>Очистить</button>
            </div>
        </div>
    );
};

export default PasswordForm;

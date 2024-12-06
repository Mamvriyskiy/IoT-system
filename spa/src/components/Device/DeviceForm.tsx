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
        <div>
            <h1>Умный дом</h1>
            <h2>Устройства</h2>
            <table>
                <thead>
                    <tr>
                        <th>Устройство</th>
                        <th>Статус</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Пылесос</td>
                        <td>On</td>
                        <td>
                            <button className={styles.delete}>Удалить</button>
                            <button className={styles.start}>Запустить</button>
                        </td>
                    </tr>
                    <tr>
                        <td>Чайник</td>
                        <td>Off</td>
                        <td>
                            <button className={styles.delete}>Удалить</button>
                            <button className={styles.start}>Запустить</button>
                        </td>
                    </tr>
                </tbody>
            </table>
            <div style={{ textAlign: "center" }}>
                <input type="text" placeholder="Имя устройства"/>
                <button className={styles.addButton}>Добавить</button>
            </div>

            <h2>Пользователи</h2>
            <table>
                <thead>
                    <tr>
                        <th>Пользователь</th>
                        <th>Уровень доступа</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>mamre@mail.ru</td>
                        <td>4</td>
                        <td>
                            <button className={styles.delete}>Удалить</button>
                        </td>
                    </tr>
                    <tr>
                        <td>ivan@yandex.ru</td>
                        <td>
                            <div className={styles.dropdown}>
                                <button>Выберите элементы</button>
                                <div className={styles.dropdownContent}>
                                    <label>
                                        <input type="checkbox" name="option1" value="1"/> Опция 1
                                    </label>
                                    <label>
                                        <input type="checkbox" name="option2" value="2"/> Опция 2
                                    </label>
                                    <label>
                                        <input type="checkbox" name="option3" value="3"/> Опция 3
                                    </label>
                                </div>
                            </div>
                        </td>
                        <td>
                            <button className={styles.delete}>Удалить</button>
                        </td>
                    </tr>
                </tbody>
            </table>
            <div style={{ textAlign: "center" }}>
                <input type="email" placeholder="Почта"/>
                <input type="number" placeholder="Доступ"/>
                <button className={styles.addButton}>Добавить</button>
            </div>
        </div>
    );
};

export default PasswordForm;

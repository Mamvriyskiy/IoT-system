import React, { useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import styles from "./DeviceForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

const PasswordForm: React.FC = () => {
    const [newPassword, setNewPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const [deviceCharacteristics, setDeviceCharacteristics] = useState([
        { characteristic: "Вес", value: 5, unit: "Кг" }
    ]);

    const location = useLocation();
    const { name } = location.state || {};

    const [deviceHistory, setDeviceHistory] = useState([
        { date: "02.01.2003", time: "29 мин", resource: "20вт" }
    ]);
    const navigate = useNavigate();

    // Функция для добавления новой характеристики устройства
    const addCharacteristic = (newCharacteristic: { characteristic: string; value: number; unit: string }) => {
        setDeviceCharacteristics((prev) => [...prev, newCharacteristic]);
    };

    // Функция для добавления новой записи в историю работы устройства
    const addHistory = (newHistory: { date: string; time: string; resource: string }) => {
        setDeviceHistory((prev) => [...prev, newHistory]);
    };

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
            <h1>Устройство: {name}</h1>
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
                    {deviceCharacteristics.map((item, index) => (
                        <tr key={index}>
                            <td>{item.characteristic}</td>
                            <td>{item.value}</td>
                            <td>{item.unit}</td>
                        </tr>
                    ))}
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
                    {deviceHistory.map((item, index) => (
                        <tr key={index}>
                            <td>{item.date}</td>
                            <td>{item.time}</td>
                            <td>{item.resource}</td>
                        </tr>
                    ))}
                </tbody>
            </table>

            <div style={{ textAlign: "center" }}>
                <button className={styles.addButton} onClick={() => addCharacteristic({ characteristic: "Температура", value: 22, unit: "°C" })}>
                    Добавить характеристику
                </button>
                {/* <button className={styles.addButton} onClick={() => addHistory({ date: "03.01.2003", time: "30 мин", resource: "25вт" })}>
                    Добавить историю работы
                </button> */}
            </div>
        </div>
    );
};

export default PasswordForm;

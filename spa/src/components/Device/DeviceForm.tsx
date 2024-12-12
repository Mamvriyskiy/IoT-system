import React, { useState, useEffect } from "react";
import { useNavigate, useLocation, useParams } from "react-router-dom";
import { getListHistory } from "../../api/authApi";
import styles from "./DeviceForm.module.css"; 
// import globStyles from "../../styles/global.module.css"; 

// Типы данных
type HistoryData = {
    timework: number;
    average: string;
    energy: number;
};

const PasswordForm: React.FC = () => {
    const [deviceHistory, setHistory] = useState<HistoryData[]>([]);

    const location = useLocation();
    // const { name } = location.state || {};

    // const [deviceHistory, setDeviceHistory] = useState([
    //     { date: "02.01.2003", time: "29 мин", resource: "20вт" }
    // ]);
    const navigate = useNavigate();

    // Функция для добавления новой характеристики устройства

    const { homeId } = useParams();
    const { deviceId } = useParams();

    useEffect(() => {
        console.log("HomeID: ", homeId, "deviceId: ", deviceId)
        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }

        if (!deviceId) {
            alert("Ошибка: deviceId не найден в URL");
            return;
        }
        
        const fetchClients = async () => {
            try {
                const response = await getListHistory(homeId, deviceId);
                console.log("История устройства получена:", response); // Для отладки
                // console.log(response)
                setHistory(response); // Assuming response is an array of home objects
            } catch (error) {
                console.error("Ошибка получения истории устройства:", error);
            }
        };
        fetchClients();
    }, [homeId, deviceId]);

    console.log("History: ", deviceHistory)

    return (
        <div>
            <h2>История работы устройства</h2>
            <table>
                <thead>
                    <tr>
                        <th>Дата работы</th>
                        <th>Время работы</th>
                        <th>Потраченные ресурсы</th>
                    </tr>
                </thead>
                <tbody>
                    {deviceHistory && deviceHistory.length > 0 ? (
                        deviceHistory.slice(0, 6).map((history, index) => (
                            <tr key={index}>
                                <td>
                                    {history.timework}
                                </td>

                                <td>
                                    {history.average}
                                </td>

                                <td>
                                    {history.energy}
                                </td>
                            </tr>
                        ))
                        ):(
                            <tr>
                                <td colSpan={4} style={{ textAlign: 'center' }}>
                                    История пуста
                                </td>
                            </tr>
                        )}
                </tbody>
            </table>
            <button className={`${styles.addButton} ${styles.centeredButton}`} onClick={() => navigate(`/api/homes/${homeId}`)}>
                Назад
            </button>
        </div>
    );
};

export default PasswordForm;

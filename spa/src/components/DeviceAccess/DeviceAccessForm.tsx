import React, { useState, useEffect } from "react";
import { useNavigate, useParams, Link, useLocation } from "react-router-dom";
import { addDevice, getListDevice, deleteDevice, addClient, getListClient, deleteClient } from "../../api/authApi";
import styles from "./DeviceAccessForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 


// Типы данных
type DevicesInfo = {
    TypeDevice: string;
    Status: string;
    Brand: string;
    Name: string;
};

type DevicesData = DevicesInfo & {
    ID: string;
    HomeID: string;
};

type ClientHome = {
    Home: string;
    login: string;
    email: string;
    AccessStatus: string;
    accesslevel: number;
    ID: string;
};

const DeviceAccessForm: React.FC = () => {
    const [newDeviceName, setNewDeviceName] = useState("");
    const [addClientEmail, setAddClientEmail] = useState("");
    const [devices, setDevices] = useState<DevicesData[]>([]); // Массив устройств
    const [clients, setClients] = useState<ClientHome[]>([]); // Массив пользователей
    const navigate = useNavigate();
    const { homeId } = useParams();

    const location = useLocation();
    const { nameHome } = location.state || {};

    // Эмуляция получения данных устройств и пользователей
    useEffect(() => {
        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }
        
        const fetchDevices = async () => {
            try {
                const response = await getListDevice(homeId);
                console.log("Устройства получены:", response); // Для отладки
                setDevices(response); // Assuming response is an array of home objects
            } catch (error) {
                console.error("Error fetching devices:", error);
            }
        };
        fetchDevices();
    }, [homeId]); // Делаем зависимость от homeId

    useEffect(() => {
        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }
        
        const fetchClients = async () => {
            try {
                const response = await getListClient(homeId);
                console.log("Пользователи получены:", response); // Для отладки
                setClients(response); // Assuming response is an array of home objects
            } catch (error) {
                console.error("Error fetching clients:", error);
            }
        };
        fetchClients();
    }, [homeId]);

    // Функция для преобразования статуса устройства
    const getDeviceStatus = (status: string) => {
        return status === "active" ? "on" : "off";
    };

    const handleSubmitDeleteDevice = async (event: React.MouseEvent, deviceID: string) => {
        event.preventDefault(); // Предотвращаем поведение по умолчанию

        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }

        try {
            console.log(`Удаляем устройство: ${deviceID}`);
            await deleteDevice({ deviceID, homeId });
            const response = await getListDevice(homeId);
            setDevices(response);
        } catch (error) {
            console.error("Ошибка при удалении устройства:", error);
            alert("Не удалось удалить устройство");
        }
    };

    const handleSubmitDeleteClient = async (event: React.MouseEvent, accessID: string) => {
        event.preventDefault(); // Предотвращаем поведение по умолчанию

        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }

        try {
            console.log(`Удаляем пользователя: ${accessID}`);
            await deleteClient({ accessID, homeId });
            const response = await getListClient(homeId);
            setClients(response);
        } catch (error) {
            console.error("Ошибка при удалении пользователя:", error);
            alert("Не удалось удалить пользователя");
        }
    };
    
    const handleSubmitDevice = async (event: React.FormEvent) => {
        event.preventDefault();

        if (newDeviceName.trim() === "") {
            alert("Введите имя устройства");
            return;
        }

        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }

        try {
            // После добавления устройства обновляем список устройств
            await addDevice({ name: newDeviceName, homeID: homeId });
            const response = await getListDevice(homeId);
            console.log("После добавления устройства, устройства:", response);

            // Обрабатываем данные перед сохранением в состоянии
            setDevices(response); // Обновляем состояние устройств
            setNewDeviceName(""); // Очистить поле ввода

        } catch (error) {
            alert("Ошибка добавления устройства");
        }
    };

    const handleSubmitAddClient = async (event: React.FormEvent) => {
        event.preventDefault();

        if (addClientEmail.trim() === "") {
            alert("Введите почту пользователя");
            return;
        }

        if (!homeId) {
            alert("Ошибка: homeId не найден в URL");
            return;
        }

        try {
            // После добавления устройства обновляем список устройств
            await addClient({ email: addClientEmail, homeID: homeId });
            const response = await getListClient(homeId);
            console.log("После добавления пользователя, пользователи:", response);

            // Обрабатываем данные перед сохранением в состоянии
            setClients(response); // Обновляем состояние устройств
            setAddClientEmail(""); // Очистить поле ввода

        } catch (error) {
            alert("Ошибка добавления пользователя");
        }
    };

    return (
        <div>
            <h1>Умный дом: {nameHome}</h1>
            <h2>Устройства</h2>
            <div style={{ overflowY: "auto", maxHeight: "300px" }}>
                <table>
                    <thead>
                        <tr>
                            <th>Устройство</th>
                            <th>Статус</th>
                            <th>Бренд</th>
                            <th>Действия</th>
                        </tr>
                    </thead>
                    <tbody>
                        {devices && devices.length > 0 ? (
                            devices.slice(0, 4).map((device) => (
                                <tr key={device.ID}>
                                    <td>
                                        <Link 
                                            to={{
                                                pathname: `/api/homes/${homeId}/devices/${device.ID}`,
                                            }}
                                            state={{ name: device.Name }} // Передаем state отдельно
                                            className={styles.link}
                                        >
                                            {device.Name}
                                        </Link>
                                        {/* <a href={`/api/homes/:homeId/devices/{${device.ID}}`} className={styles.link}>
                                            {device.Name}
                                        </a> */}
                                    </td>

                                    {/* Преобразуем статус */}
                                    <td>{getDeviceStatus(device.Status)}</td>

                                    <td>{device.Brand}</td>
                                    <td>
                                        <button className={styles.delete} onClick={(e) => handleSubmitDeleteDevice(e, device.ID)}>
                                            Удалить
                                        </button>
                                        <button className={styles.start}>Запустить</button>
                                    </td>
                                </tr>
                            ))
                            ):(
                                <tr>
                                    <td colSpan={4} style={{ textAlign: 'center' }}>
                                        Список устройств пуст
                                    </td>
                                </tr>
                            )}
                    </tbody>
                </table>
            </div>
            <div style={{ textAlign: "center" }}>
                <input
                    type="text"
                    placeholder="Имя устройства"
                    value={newDeviceName}
                    onChange={(e) => setNewDeviceName(e.target.value)}
                    required
                />
                <button className={styles.addButton} onClick={handleSubmitDevice}>
                    Добавить
                </button>
            </div>

            <h2>Пользователи</h2>
            <div style={{ overflowY: "auto", maxHeight: "300px" }}>
                <table>
                    <thead>
                        <tr>
                            <th>Почта</th>
                            <th>Пользователь</th>
                            <th>Список прав</th>
                            <th>Действия</th>
                        </tr>
                    </thead>
                    <tbody>
                        {clients && clients.length > 0 ? (
                            clients.slice(0, 4).map((client, index) => (
                                <tr key={index}>
                                    <td>{client.email}</td>
                                    <td>{client.login}</td>
                                    <td>
                                        {client.accesslevel === 4 ? (
                                            <div>Владелец</div>
                                        ) : (
                                            <div className={styles.dropdown}>
                                                <button>Выберите элементы</button>
                                                <div className={styles.dropdownContent}>
                                                    <label>
                                                        <input type="checkbox" name="option1" value="1" /> Добавление устройства
                                                    </label>
                                                    <label>
                                                        <input type="checkbox" name="option2" value="2" /> Удаление устройства
                                                    </label>
                                                    <label>
                                                        <input type="checkbox" name="option3" value="3" /> Добавление пользователя
                                                    </label>
                                                    <label>
                                                        <input type="checkbox" name="option4" value="4" /> Удаление пользователя
                                                    </label>
                                                </div>
                                            </div>
                                        )}
                                    </td>
                                    <td>
                                        {client.accesslevel < 4 ? (
                                            <button className={styles.delete} onClick={(e) => handleSubmitDeleteClient(e, client.ID)}>
                                                Удалить
                                            </button>
                                        ) : (
                                            <label></label>
                                        )}
                                    </td>
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan={4} style={{ textAlign: 'center' }}>
                                    Список пользователей пуст
                                </td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
            <div style={{ textAlign: "center" }}>
                <input
                    type="email"
                    placeholder="Почта"
                    value={addClientEmail}
                    onChange={(e) => setAddClientEmail(e.target.value)}
                    required
                />
                <button className={styles.addButton} onClick={handleSubmitAddClient}>
                    Добавить
                </button>
            </div>
        </div>
    );
};

export default DeviceAccessForm;

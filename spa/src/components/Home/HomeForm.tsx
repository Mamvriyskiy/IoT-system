import React, { useState, useEffect } from "react";
import { useNavigate, Link } from "react-router-dom";
import { getListHome, addHome, deleteHome } from "../../api/authApi";
import styles from "./HomeForm.module.css"; 
import globStyles from "../../styles/global.module.css"; 

type Home = {
    name: string;
    latitude: number;
    longitude: number;
    ID: string;
};

const HomeForm: React.FC = () => {
    const [homes, setHomes] = useState<Home[]>([]);
    const [newHomeName, setNewHomeName] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const fetchHomes = async () => {
            try {
                const response = await getListHome();
                setHomes(response); // Assuming response is an array of home objects
            } catch (error) {
                console.error("Error fetching homes:", error);
            }
        };
        fetchHomes();
    }, []);

    const handleSubmitDelete = async (event: React.MouseEvent, homeId: string) => {
        event.preventDefault(); // Предотвращаем поведение по умолчанию
        try {
            console.log(`Удаляем дом: ${homeId}`);
            await deleteHome({ homeId });
            const response = await getListHome();
            setHomes(response);
        } catch (error) {
            console.error("Ошибка при удалении дома:", error);
            alert("Не удалось удалить дом");
        }
    };

    // Handle form submission for adding a new home
    const handleSubmitHome = async (event: React.FormEvent) => {
        //event.preventDefault();

        if (newHomeName.trim() === "") {
            alert("Введите имя дома");
            return;
        }

        try {
            // Call the addHome function to add a new home
            console.log(newHomeName)
            await addHome({ name: newHomeName });
            // Re-fetch the homes after adding a new one
            const response = await getListHome();
            setHomes(response);
            setNewHomeName(""); // Clear the input field
        } catch (error) {
            alert("Ошибка добавления дома");
        }
    };

    return (
        <div>
            <h1>Умные дом</h1>
            <table>
                <thead>
                    <tr>
                        <th>Имя</th>
                        <th>Широта</th>
                        <th>Долгота</th>
                        <th>Действие</th>
                    </tr>
                </thead>
                <tbody>
                    {homes && homes.length > 0 ? (
                        homes.slice(0, 5).map((home, index) => (
                            <tr key={index}>
                                <td>
                                    <Link 
                                            to={{
                                                pathname: `/api/homes/${home.ID}`,
                                            }}
                                            state={{ nameHome: home.name }}
                                            className={styles.link}
                                        >
                                            {home.name}
                                    </Link>
                                    {/* <a href={`/api/homes/${home.ID}`} className={styles.link}>
                                        {home.name}
                                    </a> */}
                                </td>
                                <td>{home.latitude}</td>
                                <td>{home.longitude}</td>
                                <td>
                                    <button className={styles.delete} onClick={(e) => handleSubmitDelete(e, home.ID)}>
                                        Удалить
                                    </button>
                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan={4} style={{ textAlign: 'center' }}>
                                Список домов пуст
                            </td>
                        </tr>
                    )}
                </tbody>  
            </table>
            <div style={{ textAlign: "center" }}>
                <input
                    type="text"
                    placeholder="Имя дома"
                    value={newHomeName}
                    onChange={(e) => setNewHomeName(e.target.value)}
                    required
                />
                <button className={styles.addButton} onClick={handleSubmitHome}>
                    Добавить
                </button>
            </div>
        </div>
    );
};

export default HomeForm;

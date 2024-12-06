import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import AuthForm from "./components/Auth/AuthForm";
import RegistrationForm from "./components/Registration/RegistrationForm";
import CodeForm from "./components/Code/CodeForm";
import Verification from "./components/Verification/VerificationForm";
import PasswordForm from "./components/Password/PasswordForm";
import DeviceAccess from "./components/DeviceAccess/DeviceAccessForm";
import Device from "./components/Device/DeviceForm";
import Home from "./components/Home/HomeForm";

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/auth/sign-in" element={<AuthForm />} />
                <Route path="/auth/sign-up" element={<RegistrationForm />} />
                <Route path="/auth/code" element={<CodeForm />} />
                <Route path="/auth/verification" element={<Verification />} />
                <Route path="/auth/password" element={<PasswordForm />} />
                <Route path="/api/homes/:homeId" element={<DeviceAccess />} />
                <Route path="/api/homes/devices" element={<Device />} />
                <Route path="/api/homes" element={<Home />} />
            </Routes>
        </Router>
    );
};

export default App;

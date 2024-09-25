import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Button from "@/components/ui/button";
import api from "@/services/api";
import User from "@/models/User";

type FormData = {
  email: string;
  password: string;
  confirmPassword: string;
}





const SignIn = () => {
  const [response, setResponse] = useState({
    isRegister: false,
  })

  const [form, setForm] = useState<FormData>({
    email: "",
    password: "",
    confirmPassword: "",
  })

  const [passwordError, setPasswordError] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prev) => ({
      ...prev,
      [e.target.name]: e.target.value
    }))
  }


  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();


    if (form.password !== form.confirmPassword) {
      setPasswordError("Passwords do not match.");
      return;
    }
    setPasswordError("")

    const user: User = {
      email: form.email,
      password: form.password
    }

    let res = await api.user.register(user);
    if (res) {
      setResponse((prev: any) => ({
        ...prev,
        isRegister: true
      }));
    }
  }

  const IsValidRegister = ({ response }: { response: any }) => {
    const navigate = useNavigate()
    useEffect(() => {
      if (response.isRegister) {
        navigate("/login")
      }
    }, [response.isRegister]);
    return null;
  };




  return (
    <div className="flex flex-col justify-center items-center h-screen">
      <form onSubmit={handleSubmit} className="w-full max-w-sm">
        <div className="mb-4">
          <label htmlFor="email" className="block text-gray-700 font-bold mb-2">
            Email
          </label>
          <input
            type="email"
            id="email"
            name="email"
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-gray-700 font-bold mb-2">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="confirmPassword" className="block text-gray-700 font-bold mb-2">
            Confirm Password
          </label>
          <input
            type="password"
            id="confirmPassword"
            name="confirmPassword"
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
          <p className="text-red-500">{passwordError}</p>
        </div>
        <Button type="submit">Sign Up</Button>
        <p className="text-center mt-2">
          Already have an account? <a href="/login" className="text-blue-500">Login</a>
        </p>
      </form>
      <IsValidRegister response={response} />
    </div>
  )


}

export default SignIn;

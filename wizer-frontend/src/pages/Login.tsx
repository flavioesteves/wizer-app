import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Button from "@/components/ui/button";
import api from "@/services/api";
import User from "@/models/User";


type FormData = {
  email: string;
  password: string;
}

const Login = () => {
  const [response, setResponse] = useState({
    isValid: false,
    token: ""
  })
  const [form, setForm] = useState<FormData>({
    email: "",
    password: "",
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prev) => ({
      ...prev,
      [e.target.name]: e.target.value
    }))
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const user: User = {
      email: form.email,
      password: form.password
    }
    let res = await api.auth.login(user)
    if (res) {
      setResponse((prev) => ({
        ...prev,
        isValid: res.isValid,
        token: res.token
      }))
    }
  }


  const IsValidLogin = ({ response }: { response: any }) => {
    const navigate = useNavigate();
    useEffect(() => {
      if (response.isValid) {
        navigate("/");
      }
    }, [response.isValid]);
    return null;
  };


  return (
    <div className="flex justify-center items-center h-screen">
      <div className="w-96 p-6 shadow-lg bg-white rounded-md justify-center items-center flex flex-col">
        <img className="flex-none"
          src="/wizer1.svg"
          alt="wizer-logo"
          style={{ objectFit: 'cover', maxHeight: '12rem', }}
        />

        <h1 className="text-xl font-bold">Welcome</h1>

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
          <Button className="w-full bg-black hover:bg-blue-700" type="submit">Log In</Button>
          <p className="text-center mt-2">
            Don't have an account? <a href="/signin" className="text-blue-500">Sign In</a>
          </p>
        </form>

        <IsValidLogin response={response} />
      </div>
    </div>
  )

}

export default Login;

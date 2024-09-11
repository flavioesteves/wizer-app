import React, { useState } from "react";
import Button from "@/components/ui/button";
import api from "@/services/api";
import User from "@/models/User";


type FormData = {
  email: string;
  password: string;
}

const LogIn = () => {
  const [response, setResponse] = useState("")
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
    let res = api.auth.login(user)
    if (res !== null) {
      setResponse(res.toString())
      console.log(response)
    }


  }

  return (
    <div className="flex justify-center items-center h-screen">
      <form onSubmit={handleSubmit} className="w-full max-w-sm">
        <div className="mb-4">
          <label htmlFor="mb-4" className="block text-gray-700 font-bold mb-2">
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
          <label htmlFor="mb-4" className="block text-gray-700 font-bold mb-2">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
          <Button type="submit">Log In</Button>
          <p className="text-center mt-2">
            Don't have an account? <a href="/signin" className="text-blue-500">Sign In</a>
          </p>
        </div>
      </form>

    </div>
  )

}

export default LogIn;

"use client"
import React, {useEffect, useState} from 'react';
import axios from "axios";

export default function Public() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    axios({
      url: process.env.NEXT_PUBLIC_API_ENDPOINT + "/public",
      method: "GET",
  })
      .then((res) => {
        setMessage(res.data.message)
      })
      .catch((err) => {
        setMessage(err.message)
      });
  }, [])

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <p>{message}</p>
      </div>
    </main>
  )
}

import { useEffect, useState } from "react"

import "./style.css"

import { SecureStorage } from "@plasmohq/storage/secure"

const API_URL = process.env.PLASMO_PUBLIC_API_URL

interface Response {
  code: number
  success: boolean
  message: string
  data: any
}

const storage = new SecureStorage()

function IndexPopup() {
  const [codeInput, setCodeInput] = useState("")
  const [codes, setCodes] = useState<string[]>([])
  const [submitting, setSubmitting] = useState(false)
  const [copiedCode, setCopiedCode] = useState<string | null>(null)
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [loading, setLoading] = useState(true)

  const addCouponCode = async () => {
    setSubmitting(true)
    try {
      const session = await storage.get("session")
      const [tab] = await chrome.tabs.query({
        active: true,
        currentWindow: true
      })

      const response = await fetch(API_URL + "/v1/coupons", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${session}`,
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          domain: new URL(tab.url).hostname,
          code: codeInput
        })
      })

      const data = (await response.json()) as any

      if (data.success) {
        fetchCouponCodes()
        setCodeInput("")
        return
      }

      alert(data.message)
    } catch (error) {
      alert(error)
      console.log(error)
    } finally {
      setSubmitting(false)
    }
  }

  const fetchCouponCodes = async () => {
    try {
      const session = await storage.get("session")
      const [tab] = await chrome.tabs.query({
        active: true,
        currentWindow: true
      })

      const response = await fetch(
        API_URL + "/v1/coupons/" + new URL(tab.url).hostname,
        {
          headers: {
            Authorization: `Bearer ${session}`
          }
        }
      )

      const data = (await response.json()) as Response

      if (data.success) {
        setCodes(data.data as string[])
        return
      }

      alert(data.message)
    } catch (error) {
      alert(error)
      console.log(error)
    }
  }

  const emailRegister = async () => {
    try {
      const response = await fetch(API_URL + "/v1/auth/email/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email,
          password
        })
      })

      const data = (await response.json()) as Response

      if (data.success) {
        try {
          await storage.set("session", data.data as string)
          setIsAuthenticated(true)
          setEmail("")
          setPassword("")
        } catch (storageError) {
          console.error("Failed to store session:", storageError)
          alert(
            "Registration successful but failed to save session. Please try logging in."
          )
        }
      } else {
        alert(data.message)
      }
    } catch (error) {
      console.error("Registration failed:", error)
      alert("Registration failed. Please check your connection and try again.")
    }
  }

  const emailLogin = async () => {
    try {
      const response = await fetch(API_URL + "/v1/auth/email/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email,
          password
        })
      })

      const data = (await response.json()) as Response

      if (data.success) {
        try {
          await storage.set("session", data.data as string)
          setIsAuthenticated(true)
          setEmail("")
          setPassword("")
        } catch (storageError) {
          console.error("Failed to store session:", storageError)
          alert(
            "Login successful but failed to save session. Please try again."
          )
        }
      } else {
        alert(data.message)
      }
    } catch (error) {
      console.error("Login failed:", error)
      alert("Login failed. Please check your connection and try again.")
    }
  }

  const logout = async () => {
    await storage.remove("session")
    setIsAuthenticated(false)
  }

  const copyToClipboard = (code: string) => {
    navigator.clipboard.writeText(code).then(
      () => {
        setCopiedCode(code)
        setTimeout(() => {
          setCopiedCode(null)
        }, 2000)
      },
      (err) => {
        console.error("Could not copy text: ", err)
      }
    )
  }

  useEffect(() => {
    const init = async () => {
      setLoading(true)
      await storage.setPassword("sugar-password")

      const session = await storage.get("session")
      setIsAuthenticated(!!session)

      if (session) {
        await fetchCouponCodes()
      }

      setLoading(false)
    }

    init()
  }, [])

  useEffect(() => {
    if (isAuthenticated) fetchCouponCodes()
  }, [isAuthenticated])

  if (loading) {
    return (
      <div className="p-4 w-64 h-fit flex flex-col gap-4 bg-neutral-900">
        <h2 className="text-sm font-bold text-white">Sugar Baby</h2>
        <p className="text-white">Loading...</p>
      </div>
    )
  }

  return (
    <div className="p-4 w-64 h-fit flex flex-col gap-4 bg-neutral-900">
      <div className="flex flex-col gap-2">
        <div className="flex flex-row justify-between items-center">
          <h2 className="text-sm font-bold text-white">Sugar Baby</h2>
          {isAuthenticated ? (
            <button
              onClick={logout}
              className="text-sm text-white hover:text-whites/75">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="12"
                height="12"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
                className="lucide lucide-log-out">
                <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
                <polyline points="16 17 21 12 16 7" />
                <line x1="21" x2="9" y1="12" y2="12" />
              </svg>
            </button>
          ) : null}
        </div>

        {!isAuthenticated ? (
          <div className="flex flex-col gap-2">
            <input
              type="email"
              className="p-2 border rounded bg-neutral-800 border-pink-300 text-white"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <input
              type="password"
              className="p-2 border rounded bg-neutral-800 border-pink-300 text-white"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <div className="flex gap-2">
              <button
                onClick={emailLogin}
                className="bg-pink-400 py-2 px-4 rounded text-white flex-1">
                Login
              </button>
              <button
                onClick={emailRegister}
                className="bg-pink-400 py-2 px-4 rounded text-white flex-1">
                Register
              </button>
            </div>
          </div>
        ) : (
          <>
            <div className="flex flex-row gap-2">
              <input
                className="p-2 border rounded bg-neutral-800 border-pink-300 text-white"
                placeholder="CODE"
                onChange={(e) => setCodeInput(e.target.value)}
                value={codeInput}
              />
              <button
                disabled={submitting}
                onClick={addCouponCode}
                className="bg-pink-400 py-2 px-4 rounded text-white disabled:opacity-75">
                Submit
              </button>
            </div>
          </>
        )}
      </div>

      {isAuthenticated && (
        <div className="gap-1 flex flex-col">
          <p className="text-white">
            Codes Found: ({codes ? codes.length : 0})
          </p>
          {codes
            ? codes.map((code) => (
                <div
                  key={code}
                  className="flex flex-row justify-between items-center p-2 rounded border border-pink-300 bg-neutral-950 relative">
                  <p className="text-white">{code}</p>
                  <button
                    onClick={() => copyToClipboard(code)}
                    className="relative">
                    {copiedCode === code ? (
                      <span className="absolute -top-8 -left-16 bg-gray-800 text-white text-xs py-1 px-2 rounded animate-fade-in-out">
                        Copied!
                      </span>
                    ) : null}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="12"
                      height="12"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="white"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      className="lucide lucide-copy">
                      <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
                      <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
                    </svg>
                  </button>
                </div>
              ))
            : null}
        </div>
      )}
    </div>
  )
}

export default IndexPopup

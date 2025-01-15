import { useEffect, useState } from "react"

import "./style.css"

const API_URL = process.env.PLASMO_PUBLIC_API_URL

interface Response {
  code: number
  success: boolean
  message: string
  data: any
}

function IndexPopup() {
  const [codeInput, setCodeInput] = useState("")
  const [codes, setCodes] = useState<string[]>([])

  const addCouponCode = async () => {
    try {
      const response = await fetch(API_URL + "/v1/coupons/", {
        method: "POST",
        headers: {
          Authorization: "Bearer e462de5c-5109-400c-aa22-50225d20d508a",
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          domain: "www.myplots.co",
          code: codeInput
        })
      })

      const data = (await response.json()) as Response

      if (data.success) {
        fetchCouponCodes()
        return
      }

      alert(data.message)
    } catch (error) {
      console.log(error)
    }
  }

  const fetchCouponCodes = async () => {
    try {
      const [tab] = await chrome.tabs.query({
        active: true,
        currentWindow: true
      })

      const response = await fetch(
        API_URL + "/v1/coupons/" + new URL(tab.url).hostname,
        {
          headers: {
            Authorization: "Bearer e462de5c-5109-400c-aa22-50225d20d508a"
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

  useEffect(() => {
    fetchCouponCodes()
  }, [])

  return (
    <div className="p-4 w-64 flex flex-col gap-2">
      <div className="flex flex-col gap-1">
        <h2 className="text-sm font-medium">Sugar Baby</h2>
        <div className="flex flex-row gap-2">
          <input
            className="p-2 border rounded"
            placeholder="CODE"
            onChange={(e) => setCodeInput(e.target.value)}
            value={codeInput}
          />
          <button
            onClick={addCouponCode}
            className="bg-pink-400 py-2 px-4 rounded text-white">
            Submit
          </button>
        </div>
      </div>
      <div>
        <p>Codes for this site ({codes.length})</p>
        {codes.length > 0 ? (
          <>
            {codes.map((code) => (
              <p>{code}</p>
            ))}
          </>
        ) : null}
      </div>
    </div>
  )
}

export default IndexPopup

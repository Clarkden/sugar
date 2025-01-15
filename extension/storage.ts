import { SecureStorage } from "@plasmohq/storage/secure"

const secureStorage = new SecureStorage()

// // Initialize the secure storage with a password
// await secureStorage.setPassword("sugar-secure-storage")

export const storage = {
  getSession: async (): Promise<string | undefined> => {
    try {
      return await secureStorage.get("session")
    } catch (error) {
      console.error("Error getting session:", error)
      return undefined
    }
  },

  setSession: async (session: string): Promise<void> => {
    try {
      await secureStorage.set("session", session)
    } catch (error) {
      console.error("Error setting session:", error)
      throw error
    }
  },

  clearSession: async (): Promise<void> => {
    try {
      await secureStorage.remove("session")
    } catch (error) {
      console.error("Error clearing session:", error)
      throw error
    }
  }
} 
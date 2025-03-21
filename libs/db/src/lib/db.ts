import { PrismaClient } from "@prisma/client"

export function db() {
  const prisma = new PrismaClient()

  return prisma
}

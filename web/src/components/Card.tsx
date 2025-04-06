"use client";

import { motion } from "framer-motion";

export interface CardProps {
  title: string;
  description: string;
}

export const Card: React.FC<CardProps> = ({ title, description }) => {
  return (
    <motion.div
      animate={{ scale: 2, transition: { duration: 1 } }}
      className="p-4 bg-white rounded-2xl shadow-lg hover:shadow-xl transition"
    >
      <h3 className="text-xl font-semibold">{title}</h3>
      <p className="text-gray-600">{description}</p>
    </motion.div>
  );
};

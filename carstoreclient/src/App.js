import React, { useState } from "react";
import "./App.css";

const initialBrands = [
  { id: "1", name: "BMW", description: "German automotive company" },
  { id: "2", name: "Toyota", description: "Japanese automotive company" },
];

const initialCars = [
  { id: "1", name: "X5", brand_id: "1", year: 2020, color: "Black", price: 60000 },
  { id: "2", name: "Corolla", brand_id: "2", year: 2021, color: "White", price: 20000 },
];

function App() {
  const [brands, setBrands] = useState(initialBrands);
  const [cars, setCars] = useState(initialCars);

  const [newBrand, setNewBrand] = useState({ name: "", description: "" });
  const [newCar, setNewCar] = useState({ name: "", brand_id: "", year: "", color: "", price: "" });

  // Add CarBrand
  const addBrand = (e) => {
    e.preventDefault();
    if (!newBrand.name) return;
    setBrands([
      ...brands,
      {
        id: (brands.length + 1).toString(),
        name: newBrand.name,
        description: newBrand.description,
      },
    ]);
    setNewBrand({ name: "", description: "" });
  };

  // Add Car
  const addCar = (e) => {
    e.preventDefault();
    if (!newCar.name || !newCar.brand_id) return;
    setCars([
      ...cars,
      {
        id: (cars.length + 1).toString(),
        ...newCar,
        year: parseInt(newCar.year),
        price: parseFloat(newCar.price),
      },
    ]);
    setNewCar({ name: "", brand_id: "", year: "", color: "", price: "" });
  };

  const deleteBrand = (id) => {
    setBrands(brands.filter((b) => b.id !== id));
    setCars(cars.filter((c) => c.brand_id !== id));
  };

  const deleteCar = (id) => setCars(cars.filter((c) => c.id !== id));

  const getBrandName = (brand_id) =>
      brands.find((b) => b.id === brand_id)?.name || "Unknown";

  return (
      <div className="container">
        <h1>Car Store Dashboard</h1>

        <section className="section">
          <div className="section-header">
            <h2>Car Brands</h2>
            <form onSubmit={addBrand} className="inline-form">
              <input
                  className="input"
                  placeholder="Brand Name"
                  value={newBrand.name}
                  onChange={(e) => setNewBrand({ ...newBrand, name: e.target.value })}
              />
              <input
                  className="input"
                  placeholder="Description"
                  value={newBrand.description}
                  onChange={(e) => setNewBrand({ ...newBrand, description: e.target.value })}
              />
              <button className="btn primary" type="submit">Add Brand</button>
            </form>
          </div>
          <ul className="brand-list">
            {brands.map((b) => (
                <li key={b.id}>
                  <span className="brand-name">{b.name}</span>
                  <span className="brand-desc">{b.description}</span>
                  <button className="btn danger" onClick={() => deleteBrand(b.id)}>
                    Delete
                  </button>
                </li>
            ))}
          </ul>
        </section>

        <section className="section">
          <div className="section-header">
            <h2>Cars</h2>
            <form onSubmit={addCar} className="inline-form">
              <input
                  className="input"
                  placeholder="Car Name"
                  value={newCar.name}
                  onChange={(e) => setNewCar({ ...newCar, name: e.target.value })}
              />
              <select
                  className="input"
                  value={newCar.brand_id}
                  onChange={(e) => setNewCar({ ...newCar, brand_id: e.target.value })}
              >
                <option value="">Select Brand</option>
                {brands.map((b) => (
                    <option key={b.id} value={b.id}>
                      {b.name}
                    </option>
                ))}
              </select>
              <input
                  className="input"
                  placeholder="Year"
                  type="number"
                  value={newCar.year}
                  onChange={(e) => setNewCar({ ...newCar, year: e.target.value })}
              />
              <input
                  className="input"
                  placeholder="Color"
                  value={newCar.color}
                  onChange={(e) => setNewCar({ ...newCar, color: e.target.value })}
              />
              <input
                  className="input"
                  placeholder="Price"
                  type="number"
                  value={newCar.price}
                  onChange={(e) => setNewCar({ ...newCar, price: e.target.value })}
              />
              <button className="btn primary" type="submit">Add Car</button>
            </form>
          </div>
          <ul className="car-list">
            {cars.map((car) => (
                <li key={car.id}>
                  <span className="car-name">{car.name}</span>
                  <span className="car-brand">({getBrandName(car.brand_id)})</span>
                  <span className="car-info">
                {car.year} - {car.color} - ${car.price}
              </span>
                  <button className="btn danger" onClick={() => deleteCar(car.id)}>
                    Delete
                  </button>
                </li>
            ))}
          </ul>
        </section>
      </div>
  );
}

export default App;

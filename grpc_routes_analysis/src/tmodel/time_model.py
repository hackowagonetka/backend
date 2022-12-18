# Imports

import numpy as np
from xgboost import XGBRegressor
from abc import ABC


class Tmodel(ABC):
    def __init__(self) -> None:
        self.model = XGBRegressor() # Model initialization
        self.model.load_model('./grpc_routes_analysis/src/tmodel/models/m_v_0_final.model')             # Load model configs
        self.fact_time = 0               # Время фактического прибытия от пользователя
        self.penalty = 0                 # Штраф показателей модели

    def predict(self, params,  fact_time = 0):
        self.fact_time = fact_time
        result = self.model.predict([list(params)])
        corr_time = result[0] - self.fact_time
        self.penalty = (1 / len(params) * np.sqrt(sum(params))) / corr_time  # Штраф основан на среднеквадратичном отклонении

        result = result - (result * self.penalty)

        return result

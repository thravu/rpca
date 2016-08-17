package rpca

import "github.com/gonum/matrix/mat64"

type rpcaTestCase struct {
	skip        bool
	timeSeries  rPCAable
	options     *rpcaConfig
	description string
	expected    decomposedMatrix
}

var rpcaTestCases = []rpcaTestCase{
	{
		description: "auto diff and scaling",
		timeSeries: buildMatrix([]float64{
			7714, 6614, 5860, 5039, 3431, 4531, 6063, 5692, 5715, 5951, 5373, 3704,
			4332, 6011, 5898, 5215, 5454, 4568, 3502, 4460, 6740, 6152, 5678, 5187,
			4602, 3447, 7791, 5403, 4782, 4091, 3313, 3552, 3160, 3701, 5650, 5337,
			4954, 8884, 11007, 3878, 3971, 11096, 6876, 4834, 4753, 4112, 3394, 3577,
			5391, 4906, 5108, 4327, 3770, 2671, 2882, 3832, 3823, 3496, 2791, 2761,
			2928, 3277, 3560, 7700, 3715, 2760, 2762, 2842, 3169, 4366, 4575, 4431,
			4466, 3920, 3129, 3538, 4605, 4653, 5067, 4866, 4151, 3304, 3271, 4346,
			5157, 4747, 6334, 4800, 3014, 3577, 5305, 5762, 5361, 7177, 5144, 3742,
			4116, 5094, 5413, 5118, 5268, 4446, 3109, 3316, 4651, 4933, 8262, 5090,
			4372, 3234, 3290, 4583, 5851, 5422, 4908, 4319, 3157, 3479, 5120, 5055,
			4743, 5563, 4724, 3124, 4500, 5697, 5003, 4644, 5802, 4210, 3347, 3546,
			4633, 4636, 4607, 4388, 3812, 2810, 2969, 4359}, 7),
		options: &rpcaConfig{
			frequency: 7,
			autodiff:  true,
			forcediff: false,
			scale:     true,
			lPenalty:  1.0,
			sPenalty:  0.31305,
			verbose:   false,
		},
		expected: decomposedMatrix{
			L: mat64.NewDense(7, 20, []float64{
				0.6797, -253.29, 146.41, -436.48, -778.54, 285.37, 925.92, -0.45,
				-242.82, 138.62, -417.64, -744.08, 271.24, 882.54, 3.20, -276.75,
				163.83, -478.67, -855.72, 317.01, 1023.07, -21.71, -44.92, -8.39,
				-61.67, -92.93, 4.31, 62.85, -8.23, -170.43, 84.84, -287.42, -505.88,
				173.59, 582.68, 4.68, -290.47, 174.02, -503.36, -900.87, 335.52,
				1079.91, -3.38, -215.49, 118.32, -368.49, -654.17, 234.38, 769.36,
				-10.93, -145.26, 66.15, -242.15, -423.07, 139.65, 478.44, -22.41,
				-38.44, -13.21, -50.00, -71.59, -4.44, 35.99, -14.13, -115.50, 44.04,
				-188.62, -325.15, 99.51, 355.18, -9.72, -156.55, 74.53, -262.45,
				-460.20, 154.87, 525.19, -11.06, -144.04, 65.24, -239.97, -419.07,
				138.01, 473.41, 5.65, -299.55, 180.77, -519.69, -930.75, 347.77,
				1117.52, -1.35, -234.36, 132.34, -402.42, -716.24, 259.83, 847.50,
				-4.22, -207.72, 112.55, -354.50, -628.57, 223.89, 737.14, -10.39,
				-150.31, 69.90, -251.23, -439.68, 146.46, 499.36, -3.94, -210.29,
				114.46, -359.14, -637.06, 227.37, 747.82, -0.35, -243.68, 139.26,
				-419.18, -746.90, 272.40, 886.09, -5.05, -200.01, 106.82, -340.64,
				-603.23, 213.50, 705.24, -7.40, -178.10, 90.54, -301.22, -531.12,
				183.94, 614.45}),
			S: mat64.NewDense(7, 20, []float64{
				0.0, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 2983.92, -1095.08, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 10.54, 0.00, 0.00, 2400.20, 1270.58, -4872.36, 0.00,
				4689.32, -2860.84, -470.73, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 2798.35, -2513.73, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 50.46, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 327.89,
				-274.80, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 2123.54, -1886.12, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
				0.00, 0.00, 0.00}),
			E: mat64.NewDense(7, 20, []float64{
				-0.6797, -846.71, -900.41, -384.52, -829.46, 814.63, 606.08, -370.55,
				265.82, 97.38, -160.36, -924.92, 356.76, 796.46, -116.20, -406.25,
				75.17, -407.33, -210.28, 640.99, 1256.93, -566.29, -429.08, -482.61,
				-523.33, -1062.07, 1355.77, -1355.77, -612.77, -520.57, -862.84,
				526.42, 113.88, 367.41, 1355.77, -317.68, -92.53, 1355.77, 1355.77,
				-1355.77, -242.52, 1355.77, -1355.77, -1355.77, -199.32, -272.51,
				-63.83, -51.38, 1044.64, -474.07, 347.26, -847.15, -314.85, -675.93,
				71.35, 471.56, 13.41, -288.56, -691.79, 20.00, 238.59, 353.44,
				247.01, 1355.77, -1355.77, -999.04, 190.62, 405.15, 227.49, 841.82,
				218.72, 12.55, -39.53, -283.55, -330.80, 254.13, 541.81, 59.06,
				558.04, -266.24, -475.03, -427.93, -171.01, 601.59, 805.35, -110.45,
				1355.77, -1014.31, -855.25, 215.23, 610.48, 458.35, -166.64, 1355.77,
				-1355.77, -685.76, 114.17, 130.50, 323.22, -87.28, 37.45, -467.50,
				-708.43, -16.89, 597.86, 292.39, 1355.77, -1355.77, -466.77, -698.32,
				-90.46, 793.64, 1271.94, -218.71, -628.46, -229.86, -524.94, 94.63,
				893.18, -64.65, -68.32, 680.74, -419.82, -853.10, 1103.60, 310.91,
				-688.95, -158.99, 1051.18, -1251.36, -259.77, -14.50, 381.76, 10.40,
				149.10, -309.54, -274.78, -470.88, -24.94, 775.55}),
			converged:  true,
			iterations: 53},
	},
	{
		description: "no diff and no scaling",
		timeSeries: buildMatrix([]float64{
			7714, 6614, 5860, 5039, 3431, 4531, 6063, 5692, 5715, 5951, 5373, 3704,
			4332, 6011, 5898, 5215, 5454, 4568, 3502, 4460, 6740, 6152, 5678, 5187,
			4602, 3447, 7791, 5403, 4782, 4091, 3313, 3552, 3160, 3701, 5650, 5337,
			4954, 8884, 11007, 3878, 3971, 11096, 6876, 4834, 4753, 4112, 3394, 3577,
			5391, 4906, 5108, 4327, 3770, 2671, 2882, 3832, 3823, 3496, 2791, 2761,
			2928, 3277, 3560, 7700, 3715, 2760, 2762, 2842, 3169, 4366, 4575, 4431,
			4466, 3920, 3129, 3538, 4605, 4653, 5067, 4866, 4151, 3304, 3271, 4346,
			5157, 4747, 6334, 4800, 3014, 3577, 5305, 5762, 5361, 7177, 5144, 3742,
			4116, 5094, 5413, 5118, 5268, 4446, 3109, 3316, 4651, 4933, 8262, 5090,
			4372, 3234, 3290, 4583, 5851, 5422, 4908, 4319, 3157, 3479, 5120, 5055,
			4743, 5563, 4724, 3124, 4500, 5697, 5003, 4644, 5802, 4210, 3347, 3546,
			4633, 4636, 4607, 4388, 3812, 2810, 2969, 4359}, 7),
		options: &rpcaConfig{
			frequency: 7,
			autodiff:  false,
			forcediff: false,
			scale:     false,
			lPenalty:  1.0,
			sPenalty:  0.31305,
			verbose:   false,
		},
		expected: decomposedMatrix{
			L: buildMatrix([]float64{
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893,
				0.004582893, 0.004582893, 0.004582893, 0.004582893, 0.004582893}, 7),
			S: buildMatrix([]float64{
				7713.995, 6613.995, 5859.995, 5038.995, 3430.995, 4530.995, 6062.995,
				5691.995, 5714.995, 5950.995, 5372.995, 3703.995, 4331.995,
				6010.995, 5897.995, 5214.995, 5453.995, 4567.995, 3501.995,
				4459.995, 6739.995, 6151.995, 5677.995, 5186.995, 4601.995,
				3446.995, 7790.995, 5402.995, 4781.995, 4090.995, 3312.995,
				3551.995, 3159.995, 3700.995, 5649.995, 5336.995, 4953.995,
				8883.995, 11006.995, 3877.995, 3970.995, 11095.995, 6875.995,
				4833.995, 4752.995, 4111.995, 3393.995, 3576.995, 5390.995,
				4905.995, 5107.995, 4326.995, 3769.995, 2670.995, 2881.995,
				3831.995, 3822.995, 3495.995, 2790.995, 2760.995, 2927.995,
				3276.995, 3559.995, 7699.995, 3714.995, 2759.995, 2761.995,
				2841.995, 3168.995, 4365.995, 4574.995, 4430.995, 4465.995,
				3919.995, 3128.995, 3537.995, 4604.995, 4652.995, 5066.995,
				4865.995, 4150.995, 3303.995, 3270.995, 4345.995, 5156.995,
				4746.995, 6333.995, 4799.995, 3013.995, 3576.995, 5304.995,
				5761.995, 5360.995, 7176.995, 5143.995, 3741.995, 4115.995,
				5093.995, 5412.995, 5117.995, 5267.995, 4445.995, 3108.995,
				3315.995, 4650.995, 4932.995, 8261.995, 5089.995, 4371.995,
				3233.995, 3289.995, 4582.995, 5850.995, 5421.995, 4907.995,
				4318.995, 3156.995, 3478.995, 5119.995, 5054.995, 4742.995,
				5562.995, 4723.995, 3123.995, 4499.995, 5696.995, 5002.995,
				4643.995, 5801.995, 4209.995, 3346.995, 3545.995, 4632.995,
				4635.995, 4606.995, 4387.995, 3811.995, 2809.995, 2968.995,
				4358.995}, 7),
			E: buildMatrix([]float64{
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543,
				0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543, 0.0008451543}, 7),
			converged:  true,
			iterations: 3},
	},
	{
		description: "scaling and no diff",
		timeSeries: buildMatrix([]float64{
			7714, 6614, 5860, 5039, 3431, 4531, 6063, 5692, 5715, 5951, 5373, 3704,
			4332, 6011, 5898, 5215, 5454, 4568, 3502, 4460, 6740, 6152, 5678, 5187,
			4602, 3447, 7791, 5403, 4782, 4091, 3313, 3552, 3160, 3701, 5650, 5337,
			4954, 8884, 11007, 3878, 3971, 11096, 6876, 4834, 4753, 4112, 3394, 3577,
			5391, 4906, 5108, 4327, 3770, 2671, 2882, 3832, 3823, 3496, 2791, 2761,
			2928, 3277, 3560, 7700, 3715, 2760, 2762, 2842, 3169, 4366, 4575, 4431,
			4466, 3920, 3129, 3538, 4605, 4653, 5067, 4866, 4151, 3304, 3271, 4346,
			5157, 4747, 6334, 4800, 3014, 3577, 5305, 5762, 5361, 7177, 5144, 3742,
			4116, 5094, 5413, 5118, 5268, 4446, 3109, 3316, 4651, 4933, 8262, 5090,
			4372, 3234, 3290, 4583, 5851, 5422, 4908, 4319, 3157, 3479, 5120, 5055,
			4743, 5563, 4724, 3124, 4500, 5697, 5003, 4644, 5802, 4210, 3347, 3546,
			4633, 4636, 4607, 4388, 3812, 2810, 2969, 4359}, 7),
		options: &rpcaConfig{
			frequency: 7,
			autodiff:  false,
			forcediff: false,
			scale:     true,
			lPenalty:  1.0,
			sPenalty:  0.31305,
			verbose:   false,
		},
		expected: decomposedMatrix{
			L: buildMatrix([]float64{
				5393.284, 5228.416, 5567.618, 4951.002, 3892.644, 4325.029, 5357.746,
				5247.284, 5133.820, 5428.270, 4946.881, 4110.784, 4450.194,
				5241.929, 5209.530, 5053.710, 5259.360, 4784.019, 3986.693,
				4316.527, 5139.953, 5085.409, 4996.199, 5195.596, 4847.125,
				4246.451, 4491.288, 5071.144, 4810.530, 4583.207, 4372.365,
				4156.415, 3895.255, 4026.106, 4549.060, 5334.836, 5260.851,
				5679.704, 5153.761, 4208.008, 4584.897, 5402.379, 5131.648,
				4892.984, 4921.737, 4461.226, 3745.403, 4054.304, 4935.451,
				4962.907, 4659.834, 4465.027, 4096.456, 3595.899, 3829.532,
				4641.329, 4487.905, 4230.897, 3722.330, 3730.740, 3912.571,
				3875.214, 4107.687, 4807.831, 4471.203, 4106.524, 3835.773,
				3541.667, 3699.461, 4404.212, 4870.074, 4656.349, 4511.723,
				4258.585, 3918.396, 4078.180, 4641.030, 4941.556, 4730.740,
				4646.991, 4342.237, 3902.663, 4100.670, 4734.075, 5232.601,
				5024.930, 5176.966, 4657.531, 3810.387, 4166.280, 5101.655,
				5224.083, 5100.358, 5362.119, 4892.641, 4085.669, 4415.095,
				5199.670, 5112.533, 4876.933, 4894.738, 4450.026, 3762.071,
				4059.752, 4915.543, 5093.041, 4870.931, 4891.959, 4468.746,
				3812.692, 4096.238, 4908.658, 5157.267, 4913.281, 4955.019,
				4472.700, 3719.119, 4043.374, 4960.561, 5100.364, 4952.169,
				5081.599, 4691.358, 4049.864, 4318.189, 5013.476, 5047.595,
				4852.084, 4873.890, 4498.278, 3914.969, 4166.819, 4886.325,
				4931.958, 4653.803, 4468.970, 4136.352, 3687.623, 3897.901,
				4634.928}, 7),
			S: buildMatrix([]float64{
				1102.6710, 167.5384, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 382.0019, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 2081.6669, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 1986.2503, 4635.1932,
				0.0000, 0.0000, 4475.5754, 526.3067, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 1674.1236, 0.0000, -128.4790, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				596.8359, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 2173.0237, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000,
				0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000, 0.0000}, 7),
			E: buildMatrix([]float64{
				1218.045422, 1218.045421, 292.381762, 87.997792, -461.643981,
				205.971072, 705.253725, 444.716154, 581.180080, 522.730482,
				426.118522, -406.783706, -118.193881, 769.070669, 688.470152,
				161.289747, 194.640402, -216.018913, -484.692755, 143.473412,
				1218.045421, 1066.590600, 681.800680, -8.596076, -245.124699,
				-799.450972, 1218.045408, 331.855640, -28.530408, -492.206837,
				-1059.365300, -604.415151, -735.254854, -325.105947, 1100.939737,
				2.163714, -306.850845, 1218.045433, 1218.045424, -330.007693,
				-613.896754, 1218.045426, 1218.045417, -58.984191, -168.736717,
				-349.225550, -351.403169, -477.304453, 455.548835, -56.907028,
				448.166495, -138.027037, -326.456035, -924.899496, -947.531621,
				-809.329084, -664.905201, -734.897471, -931.329959, -969.740404,
				-984.571235, -598.214465, -547.686789, 1218.045409, -756.203332,
				-1218.045430, -1073.773458, -699.667400, -530.460625, -38.211886,
				-295.073978, -225.348506, -45.723441, -338.585287, -789.395634,
				-540.180370, -36.029505, -288.556397, 336.259601, 219.008749,
				-191.236611, -598.663078, -829.670153, -388.075381, -75.600539,
				-277.930267, 1157.033858, 142.469232, -796.386866, -589.279740,
				203.345029, 537.916967, 260.642374, 1218.045428, 251.358654,
				-343.668889, -299.095404, -105.670143, 300.466770, 241.066947,
				373.261616, -4.026202, -653.070583, -743.752424, -264.542608,
				-160.040635, 1218.045416, 198.041061, -96.745646, -578.691719,
				-806.237528, -325.658054, 693.732504, 508.718577, -47.018689,
				-153.699931, -562.118759, -564.373711, 159.439461, -45.363564,
				-209.169136, 481.401343, 32.642368, -925.864273, 181.811405,
				683.524121, -44.595014, -208.084036, 928.109947, -288.277548,
				-567.969028, -620.819322, -253.325288, -295.958442, -46.803491,
				-80.970430, -324.352235, -877.622506, -928.900972, -275.928110}, 7),
			converged:  true,
			iterations: 66},
	},
}

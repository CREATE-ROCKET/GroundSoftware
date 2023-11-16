import pandas as pd
import sys
import matplotlib.pyplot as plt

def calculate_error_rates(csv_file):
    # Read the file line by line
    with open(csv_file, 'r') as file:
        lines = file.readlines()

    # Create a DataFrame from the lines
    string_data = pd.DataFrame(lines, columns=['line'])

    # 誤りをカウント（各行が指定した値を含むかどうか）
    values_to_check = [
        '219001775,0.998016,0.023623204,0.02477049,-0.05284419,',
        '219006777,0.99801594,0.023620263,0.024775878,-0.05284336,',
        '219012774,0.99801576,0.023606848,0.02477091,-0.05285531,',
        '219017777,0.998016,0.023609942,0.024775814,-0.052847195,',
        '218981959,0.9980157,0.023620343,0.024717813,-0.052875597,',
        '218984774,0.99801594,0.023613272,0.02471597,-0.052875385,',
        '218989774,0.99801564,0.023608612,0.02473001,-0.052875895,',
        '218995775,0.99801654,0.02362494,0.024767606,-0.052834805,'
    ]

    # Check if any of the specified values are present in each line
    errors = ~string_data['line'].str.strip().isin(values_to_check)

    # Display the DataFrame with errors
    print(string_data[errors])




    # 50行ごとの誤り率を計算
    error_rates_50 = []
    for i in range(0, len(errors), 50):
        error_count = errors[i:i+50].sum()
        error_rate = error_count / min(50, len(errors) - i)
        error_rates_50.append(error_rate)

    # 全行の誤り率を計算
    total_error_rate = errors.sum() / len(errors)

    return error_rates_50, total_error_rate

import os

def plot_error_rates(error_rates_50, total_error_rate, output_filename="error_plot.png"):
    # Plot the graph
    plt.plot(error_rates_50, marker='o', label='Error Rate per 50 Rows')
    plt.axhline(y=total_error_rate, color='r', linestyle='-', label='Total Error Rate')

    plt.xlabel('Block Number (Every 50 Rows)')
    plt.ylabel('Error Rate')
    plt.title('Error Rates Comparison')
    plt.legend()

    # Get the directory of the script
    script_directory = os.path.dirname(os.path.realpath(__file__))

    # Save the plot in the same directory as the script
    output_path = os.path.join(script_directory, output_filename)
    plt.savefig(output_path)

    # Show the plot
    plt.show()


def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <csv_file>")
        sys.exit(1)

    csv_file = sys.argv[1]
    error_rates_50, total_error_rate = calculate_error_rates(csv_file)

    print("50行ごとの誤り率:")
    for i, rate in enumerate(error_rates_50):
        print(f"ブロック {i+1}: {rate * 100:.2f}%")

    print(f"\n全行の誤り率: {total_error_rate * 100:.2f}%")

    plot_error_rates(error_rates_50, total_error_rate)

if __name__ == "__main__":
    main()
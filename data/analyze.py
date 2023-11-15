import pandas as pd
import sys
import matplotlib.pyplot as plt

def calculate_error_rates(csv_file):
    # CSVファイルを読み込む
    data = pd.read_csv(csv_file, header=None, on_bad_lines='skip')

    # NaN値をチェックするためにデータフレームを文字列に変換
    string_data = data.astype(str)

    # 誤りをカウント（'NaN' または 'e' を含む行）
    errors = string_data.apply(lambda x: x.str.contains('NaN') | x.str.contains('e')).any(axis=1)

    # 50行ごとの誤り率を計算
    error_rates_50 = []
    for i in range(0, len(errors), 50):
        error_count = errors[i:i+50].sum()
        error_rate = error_count / min(50, len(errors) - i)
        error_rates_50.append(error_rate)

    # 全行の誤り率を計算
    total_error_rate = errors.sum() / len(errors)

    return error_rates_50, total_error_rate

def plot_error_rates(error_rates_50, total_error_rate):
    # Plot the graph
    plt.plot(error_rates_50, marker='o', label='Error Rate per 50 Rows')
    plt.axhline(y=total_error_rate, color='r', linestyle='-', label='Total Error Rate')

    plt.xlabel('Block Number (Every 50 Rows)')
    plt.ylabel('Error Rate')
    plt.title('Error Rates Comparison')
    plt.legend()
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